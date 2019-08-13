package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func CreateArticleComment(db *sqlx.Tx, a *model.ArticleComment) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO article (body, user_id) VALUES (?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.Body, a.UserID)
}
