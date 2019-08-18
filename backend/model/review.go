package model

// Review に関する構造体
type Review struct {
	ID        int64  `db:"id" json:"id"`
	LectureID string `db:"lecture_id" json:"lecture_id"`
	Content   string `db:"content" json:"content"`
}
