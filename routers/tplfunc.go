package routers

import (
	"crypto/rand"
	"fmt"
	"ginblog/config"
	"github.com/pkg/errors"
	"html/template"
	"math/big"
	"reflect"
	"strings"
	"time"
)

func getConfig(returnType, key string, defaultValue interface{}) (value interface{}, err error) {
	returnType = strings.ToLower(returnType)
	switch returnType {
	case "string":
		value, err = config.C.String(key)
	case "bool":
		value, err = config.C.Bool(key)
	case "int":
		value, err = config.C.Int(key)
	case "int64":
		value, err = config.C.Int64(key)
	case "float":
		value, err = config.C.Float(key)
	case "diy":
		value, err = config.C.DIY(key)
	default:
		err = errors.New("Config keys must be of type String, Bool, Int, Int64, Float, or DIY")
	}
	if err != nil {
		if returnType != reflect.TypeOf(defaultValue).Kind().String() {
			err = errors.New("defaultVal type does not match returnType")
		} else {
			value, err = defaultValue, nil
		}
	} else if reflect.TypeOf(value).Kind() == reflect.String {
		if value == "" {
			if reflect.TypeOf(defaultValue).Kind() != reflect.String {
				err = errors.New("defaultVal type must be a String if the returnType is a String")
			} else {
				value = defaultValue.(string)
			}
		}
	}

	return
}

func subStr(s string, start, length int) string {
	bt := []rune(s)
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}

func str2html(raw string) template.HTML {
	return template.HTML(raw)
}

func getGravatar() string {
	i := randInt64(1, 5)
	return "/static/img/avatar/" + fmt.Sprintf("%d", i) + ".jpg"
}

func randInt64(min, max int64) int64 {
	maxInt := big.NewInt(max)
	i, _ := rand.Int(rand.Reader, maxInt)
	if i.Int64() < min {
		randInt64(min, max)
	}
	return i.Int64()
}

func getDate(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04")
}

func getDateMH(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("01-02 03:04")
}
