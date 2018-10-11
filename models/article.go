package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Article struct {
	ID       int64  `db:"id" json:"id" form:"id"`
	Title    string `db:"title" json:"title" form:"title"`
	Uri      string `db:"uri" json:"uri" form:"uri"`
	Keywords string `db:"keywords" json:"keywords" form:"keywords"`
	Summary  string `db:"summary" json:"summary" form:"summary"`
	Content  string `db:"content" json:"content" form:"content"`
	Author   string `db:"author" json:"author" form:"author"`
	Created  int64  `db:"created" json:"created" form:"created"`
	Viewnum  int    `db:"viewnum" json:"viewnum" form:"viewnum"`
	Status   int    `db:"status" json:"status" form:"status"`
}

// check Article field validate
func (art *Article) Validate() error {
	if art.Title == "" {
		return errors.New("标题不能为空")
	}
	if art.Content == "" {
		return errors.New("内容不能为空")
	}
	return nil
}

// get article
func GetArticle(db sqlx.Queryer, id int64) (Article, error) {
	var art Article
	err := sqlx.Get(db, &art, `
		select id,
		title,
		uri,
		keywords,
		summary,
		content,
		author,
		created,
		viewnum,
		status
		from article
		where id=?`, id)
	if err != nil {
		return art, err
	}
	return art, nil
}

// update article
func UpdateArticle(db sqlx.Execer, upArt Article) error {
	result, err := db.Exec(`
		update article
		set title=?,
		uri=?,
		keywords=?,
		summary=?,
		content=?,
		author=?,
		status=?
		where id=?`,
		upArt.Title,
		upArt.Uri,
		upArt.Keywords,
		upArt.Summary,
		upArt.Content,
		upArt.Author,
		upArt.Status,
		upArt.ID,
	)
	if err != nil {
		return err
	}
	if n, _ := result.RowsAffected(); n == 0 {
		return errors.New("no row update")
	}
	return nil
}

// create article
func CreateArticle(db sqlx.Execer, item *Article) (int64, error) {
	if err := item.Validate(); err != nil {
		return 0, errors.Wrap(err, "validate error")
	}
	result, err := db.Exec(`
		insert into article(title,
		uri,
		keywords,
		summary,
		content,
		author,
		created,
		viewnum,
		status)
		values(?,?,?,?,?,?,?,?,?)`,
		item.Title,
		item.Uri,
		item.Keywords,
		item.Summary,
		item.Content,
		item.Author,
		item.Created,
		item.Viewnum,
		item.Status,
	)
	if err != nil {
		return 0, errors.Wrap(err, "insert article error")
	}
	return result.LastInsertId()
}

// list article
func ListArticle(db sqlx.Queryer, page, offset int, status, title, keywords string) ([]Article, error) {
	var arts []Article
	query := `
	select id,
	title,
	uri,
	keywords,
	summary,
	content,
	author,
	created,
	viewnum,
	status
	from article
	where status=` + status
	if title != "" && keywords != "" {
		query += " and (title like '%" + title + "%'" + " or keywords like '%" + keywords + "%')"
	} else if title != "" {
		query += " and title like '%" + title + "%'"
	} else if keywords != "" {
		query += " and keywords like '%" + keywords + "%'"
	}

	query += "  limit ? offset ?"
	start := (page - 1) * offset
	err := sqlx.Select(db, &arts, query, offset, start)
	if err != nil {
		return arts, err
	}
	return arts, nil
}

// count article
func CountArticle(db sqlx.Queryer, status string, title, keywords string) (int64, error) {
	var count int64
	query := `select count(id) cnt	from article where status=` + status
	if title != "" && keywords != "" {
		query += " and (title like '%" + title + "%'" + " or keywords like '%" + keywords + "%')"
	} else if title != "" {
		query += " and title like '%" + title + "%'"
	} else if keywords != "" {
		query += " and keywords like '%" + keywords + "%'"
	}

	err := sqlx.Get(db, &count, query)
	if err != nil {
		return 0, err
	}
	return count, nil
}
