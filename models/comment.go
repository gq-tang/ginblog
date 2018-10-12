package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Comment struct {
	ID        int64  `db:"id" json:"id" form:"id"`
	ArticleID int64  `db:"article_id" json:"article_id" form:"article_id"`
	Nickname  string `db:"nickname" json:"nickname" form:"nickname"`
	Uri       string `db:"uri" json:"uri" form:"uri"`
	Content   string `db:"content" json:"content" form:"content"`
	Created   int64  `db:"created" json:"created" form:"created"`
	Status    int    `db:"status" json:"status" form:"status"`
}

// check comment field value
func (item *Comment) Validate() error {
	if item.Nickname == "" {
		return errors.New("nickname can't empty.")
	}
	if item.Content == "" {
		return errors.New("content can't empty.")
	}
	return nil
}

// create comment
func CreateComment(db sqlx.Execer, item *Comment) (int64, error) {
	err := item.Validate()
	if err != nil {
		return 0, err
	}
	result, err := db.Exec(`
		insert into comment(
			article_id,
			nickname,
			uri,
			content,
			created,
			status
		)values(?,?,?,?,?,?)`,
		item.ArticleID,
		item.Nickname,
		item.Uri,
		item.Content,
		item.Created,
		item.Status,
	)
	if err != nil {
		return 0, errors.Wrap(err, "create comment error")
	}
	return result.LastInsertId()
}

//update comment status
func UpdateComment(db sqlx.Execer, id int64, status int) error {
	result, err := db.Exec(`update comment set status=? where id=?`, status, id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("no row update.")
	}
	return nil
}

// list article's comment
func ListComment(db sqlx.Queryer, page, offset int, articleID int64, status string) ([]Comment, error) {
	var items []Comment
	if status != "0" && status != "1" {
		status = "status"
	}
	query := `select id,
		article_id,
		nickname,
		uri,
		content,
		created,
		status 
		from comment where article_id=?`
	if status != "" {
		query += " and status=" + status
	}
	query += " limit ? offset ?"
	start := (page - 1) * offset
	err := sqlx.Select(db, &items, query, articleID, offset, start)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return items, nil
}

// count article's comment
func CountComment(db sqlx.Queryer, articleID int64, status string) (int, error) {
	var count int
	if status != "0" && status != "1" {
		status = "status"
	}
	query := "select count(id) cnt from comment where article_id=? and status=" + status

	err := sqlx.Get(db, &count, query, articleID)

	return count, err
}
