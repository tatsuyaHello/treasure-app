package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func AllLecture(db *sqlx.DB) ([]model.Lecture, error) {
	a := make([]model.Lecture, 0)
	if err := db.Select(&a, `SELECT id, year, lecture_id, title, sub_title, english_title, unit, semester, location, lecture_style, teacher, overview, goal, evaluate_id, textbook, reference_url, remarks  FROM lectures`); err != nil {
		return nil, err
	}
	for i := 0; i < len(a); i++ {
		b := make([]model.Evaluate, 0)
		if err := db.Select(&b, `SELECT method, comment, percentage FROM evaluates WHERE id=?`, a[i].EvaluateID); err != nil {
			return nil, err
		}
		a[i].Evaluate = b
	}
	for i := 0; i < len(a); i++ {
		b := make([]model.Scehdule, 0)
		if err := db.Select(&b, `SELECT session FROM scehdules WHERE lecture_id=?`, a[i].LectureID); err != nil {
			return nil, err
		}
		a[i].Scehdule = b
	}
	return a, nil
}

// 検索ワードが以下のいずれかに含まれているものを、検索結果として返す
// タイトル, サブタイトル, 英語タイトル, 講師, 概要, ゴール

// FindLecture is
func FindLecture(db *sqlx.DB, title, english_title, semester, location, lecture_style, teacher, overview string, unit int64) ([]model.Lecture, error) {
	a := make([]model.Lecture, 0)
	if unit != 100 {
		if err := db.Select(&a, `SELECT * FROM lectures WHERE (unit = ? AND (title like ?) AND (english_title like ?) AND (semester like ?) AND (location like ?) AND (teacher like ?) AND (lecture_style like ?) AND (overview like ?))`, unit, "%"+title+"%", "%"+english_title+"%", "%"+semester+"%", "%"+location+"%", "%"+teacher+"%", "%"+lecture_style+"%", "%"+overview+"%"); err != nil {
			return nil, err
		}
	} else {
		if err := db.Select(&a, `SELECT * FROM lectures WHERE ((title like ?) AND (english_title like ?) AND (semester like ?) AND (location like ?) AND (teacher like ?) AND (lecture_style like ?) AND (overview like ?))`, "%"+title+"%", "%"+english_title+"%", "%"+semester+"%", "%"+location+"%", "%"+teacher+"%", "%"+lecture_style+"%", "%"+overview+"%"); err != nil {
			return nil, err
		}
	}
	for i := 0; i < len(a); i++ {
		b := make([]model.Evaluate, 0)
		if err := db.Select(&b, `SELECT method, comment, percentage FROM evaluates WHERE id=?`, a[i].EvaluateID); err != nil {
			return nil, err
		}
		a[i].Evaluate = b
	}
	for i := 0; i < len(a); i++ {
		b := make([]model.Scehdule, 0)
		if err := db.Select(&b, `SELECT session FROM scehdules WHERE lecture_id=?`, a[i].LectureID); err != nil {
			return nil, err
		}
		a[i].Scehdule = b
	}
	return a, nil
}

// ShowLecture is
func ShowLecture(db *sqlx.DB, lecture_id string) (*model.Lecture, error) {
	a := model.Lecture{}
	if err := db.Get(&a, `
		SELECT * FROM lectures WHERE lecture_id = ?
	`, lecture_id); err != nil {
		return nil, err
	}
	b := make([]model.Evaluate, 0)
	if err := db.Select(&b, `SELECT method, comment, percentage FROM evaluates WHERE id=?`, a.EvaluateID); err != nil {
		return nil, err
	}
	a.Evaluate = b

	c := make([]model.Scehdule, 0)
	if err := db.Select(&c, `SELECT session FROM scehdules WHERE lecture_id=?`, a.LectureID); err != nil {
		return nil, err
	}
	a.Scehdule = c
	return &a, nil
}

// CreateReview is
func CreateReview(db *sqlx.Tx, a *model.Review) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO reviews (lecture_id, content) VALUES (?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.LectureID, a.Content)
}

// GetReviews is
func GetReviews(db *sqlx.DB, id string) ([]model.Review, error) {
	a := make([]model.Review, 0)
	if err := db.Select(&a, `
SELECT id, lecture_id, content FROM reviews WHERE lecture_id = ?
`, id); err != nil {
		return nil, err
	}
	return a, nil
}
