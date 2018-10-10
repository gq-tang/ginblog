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
