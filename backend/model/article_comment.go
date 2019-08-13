package model

type CreateRequestArticleComment struct {
	Body string `json:"body"`
}

// あーめんどくさい
type ResponseCreateArticleComment struct {
}

type ArticleComment struct {
	ID        int64            `db:"id" json:"id"`
	Body      string           `db:"body" json:"body"`
	ArticleID int64            `db:"article_id" json:"article_id"`
	UserID    int64            `db:"user_id" json:"user_id"`
	Comment   []ArticleComment `json:"comment"`
}
