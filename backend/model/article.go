package model

type Article struct {
	ID     int64  `db:"id" json:"id"`
	Title  string `db:"title" json:"title"`
	Body   string `db:"body" json:"body"`
	UserID *int64 `db:"user_id" json:user_id`
}

type CreateArticleRequest struct {
	Title  string  `json:"title"`
	Body   string  `json:"body"`
	TagIDs []int64 `json:"tag_ids"`
}

type CreateArticleResponse struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Body   string  `json:"body"`
	UserID *int64  `json:"user_id"`
	Tags   []int64 `json:"tags"`
}
