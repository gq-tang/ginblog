package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"time"
)

type Album struct {
	ID       int64  `db:"id"`
	Title    string `db:"title"`
	Picture  string `db:"picture"`
	Keywords string `db:"keywords"`
	Summary  string `db:"summary"`
	Created  int64  `db:"created"`
	Viewnum  int    `db:"viewnum"`
	Status   int    `db:"status"`
}

func (a Album) Validate() error {
	return nil
}

// CreateAlbum creates the given Album.
func CreateAlbum(db sqlx.Queryer, item *Album) error {
	if err := item.Validate(); err != nil {
		return errors.Wrap(err, "validate error")
	}

	created := time.Now().Unix()
	err := sqlx.Get(db, &item.ID, `
		insert into album(
			title,
			picture,
			keywords,
			summary,
			created,
			viewnum,
			status
		)values(?,?,?,?,?,?,?);
		select last_insert_id() as id`,
		item.Title,
		item.Picture,
		item.Keywords,
		item.Summary,
		created,
		0,
		item.Status,
	)
	if err != nil {
		return errors.Wrap(err, "create album error")
	}
	return nil
}
