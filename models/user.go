package models

import (
	"bytes"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/pkg/errors"
	"golang.org/x/crypto/pbkdf2"
)

// user struct
type User struct {
	ID            int64  `db:"id"`
	Phone         string `db:"phone"`
	Password      string `db:"password"`
	UserProfileID *int64 `db:"user_profile_id"`
	Created       *int64 `db:"created"`
	Changed       *int64 `db:"changed"`
	Status        int    `db:"status"`
}

// UserProfile struct
type UserProfile struct {
	ID       int64  `db:"id"`
	Realname string `db:"realname"`
	Sex      int    `db:"sex"`
	Birth    string `db:"birth"`
	Email    string `db:"email"`
	Phone    string `db:"phone"`
	Address  string `db:"address"`
	Hobby    string `db:"hobby"`
	Intro    string `db:"intro"`
}

func hash(password string, saltSize int, iterations int) (string, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return "", errors.Wrap(err, "read random bytes error")
	}

	return hashWithSalt(password, salt, iterations), nil
}

func hashWithSalt(password string, salt []byte, iterations int) string {
	hash := pbkdf2.Key([]byte(password), salt, iterations, sha512.Size, sha512.New)

	var buf bytes.Buffer

	buf.WriteString("PBKDF2$")
	buf.WriteString("sha512$")
	buf.WriteString(strconv.Itoa(iterations))
	buf.WriteString("$")
	buf.WriteString(base64.StdEncoding.EncodeToString(salt))
	buf.WriteString("$")
	buf.WriteString(base64.StdEncoding.EncodeToString(hash))

	return buf.String()
}

func hashCompare(password, passwordHash string) (bool, error) {
	hashSplit := strings.Split(passwordHash, "$")
	if len(hashSplit) != 5 {
		return false, errors.Errorf("%s format error", passwordHash)
	}
	iterations, err := strconv.Atoi(hashSplit[2])
	if err != nil {
		return false, err
	}
	salt, err := base64.StdEncoding.DecodeString(hashSplit[3])
	if err != nil {
		return false, err
	}
	newHash := hashWithSalt(password, salt, iterations)
	return newHash == passwordHash, nil
}

// LoginUser
func LoginUser(db sqlx.Queryer, phone, password string) (User, error) {
	var user User
	err := sqlx.Get(db, &user, `select
		id,
		phone,
		password, 
		user_profile_id,
		created,
		changed,
		status
	 	from user 
	 	where phone=?`, phone)
	if err != nil {
		return user, errors.Wrap(err, "phone error")
	}
	result, err := hashCompare(password, user.Password)
	if err != nil {
		return user, errors.Wrap(err, "password error")
	}
	if !result {
		return user, errors.New("password error")
	}
	return user, nil
}

func GetUserProfile(db sqlx.Queryer, id int64) (UserProfile, error) {
	var pro UserProfile
	err := sqlx.Get(db, &pro, `
		select id,
		realname,
		sex,
		birth,
		email,
		phone,
		address,
		hobby,
		intro
		from user_profile
		where id=?`, id)
	if err != nil {
		return pro, err
	}
	return pro, nil
}
