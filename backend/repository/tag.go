package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func CreateArticleTag(db *sqlx.Tx, articleId int64, tagId int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO article_tag (article_id, tag_id) VALUES (?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(articleId, tagId)
}
