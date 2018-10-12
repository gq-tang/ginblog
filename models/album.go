package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Album struct {
	ID       int64  `db:"id" json:"id" form:"id"`
	Title    string `db:"title" json:"title" form:"title"`
	Picture  string `db:"picture" json:"picture" form:"picture"`
	Keywords string `db:"keywords" json:"keywords" form:"keywords"`
	Summary  string `db:"summary" json:"summary" form:"summary"`
	Created  int64  `db:"created" json:"created" form:"created"`
	Viewnum  int    `db:"viewnum" json:"viewnum" form:"viewnum"`
	Status   int    `db:"status" json:"status" form:"status"`
}

// check Album field validate
func (a *Album) Validate() error {
	return nil
}

// CreateAlbum creates the given Album.
func CreateAlbum(db sqlx.Execer, item *Album) (int64, error) {
	if err := item.Validate(); err != nil {
		return 0, errors.Wrap(err, "validate error")
	}

	result, err := db.Exec(`
		insert into album(
			title,
			picture,
			keywords,
			summary,
			created,
			viewnum,
			status
		)values(?,?,?,?,?,?,?)`,
		item.Title,
		item.Picture,
		item.Keywords,
		item.Summary,
		item.Created,
		item.Viewnum,
		item.Status,
	)
	if err != nil {
		return 0, errors.Wrap(err, "create album error")
	}
	return result.LastInsertId()
}

// GetAlbum
func GetAlbum(db sqlx.Queryer, id int64) (Album, error) {
	var album Album
	err := sqlx.Get(db, &album, `
		select id,
		title,
		picture,
		keywords,
		summary,
		created,
		viewnum,
		status
		from album
		where id=?`, id)
	if err != nil {
		return album, err
	}
	return album, nil
}

// UpdateAlbum
func UpdateAlbum(db sqlx.Execer, item *Album) error {
	_, err := db.Exec(`
		update album
		set title=?,
		summary=?,
		status=?
		where id=?`,
		item.Title,
		item.Summary,
		item.Status,
		item.ID,
	)
	return err
}

// ListAlbum
func ListAlbum(db sqlx.Queryer, page, offset int, title, keywords, status string) ([]Album, error) {
	var items []Album
	if status != "0" && status != "1" {
		status = "status"
	}
	query := `select id,
	title,
	picture,
	keywords,
	summary,
	created,
	viewnum,
	status
	from album
	where status=` + status
	if title != "" && keywords != "" {
		query += " and (title like '%" + title + "%' or keywords like '%" + keywords + "%')"
	} else if title != "" {
		query += " and title like '%" + title + "%'"
	} else if keywords != "" {
		query += " and keywords like '%" + keywords + "%'"
	}
	query += " limit ? offset ?"
	start := (page - 1) * offset
	err := sqlx.Select(db, &items, query, offset, start)
	if err != nil {
		return nil, err
	}
	return items, nil
}

// CountAlbum
func CountAlbum(db sqlx.Queryer, title, keywords, status string) (int, error) {
	var count int
	if status != "0" && status != "1" {
		status = "status"
	}
	query := "select count(id) cnt from album where status=" + status
	if title != "" && keywords != "" {
		query += " and (title like '%" + title + "%' or keywords like '%" + keywords + "%')"
	} else if title != "" {
		query += " and title like '%" + title + "%'"
	} else if keywords != "" {
		query += " and keywords like '%" + keywords + "%'"
	}
	err := sqlx.Get(db, &count, query)
	return count, err
}
