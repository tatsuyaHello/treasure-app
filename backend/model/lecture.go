package model

// フロントに返してやるモデル
type Lecture struct {
	ID           int64  `db:"id" json:"id"`
	Year         string `db:"year" json:"year"`
	LectureID    string `db:"lecture_id" json:"lecture_id"`
	Title        string `db:"title" json:"title"`
	SubTitle     string `db:"sub_title" json:"sub_title"`
	EnglishTitle string `db:"english_title" json:"english_title"`
	Unit         int64  `db:"unit" json:"unit"`
	Semester     string `db:"semester" json:"semester"`
	Location     string `db:"location" json:"location"`
	LectureStyle string `db:"lecture_style" json:"lecture_style"`
	Teacher      string `db:"teacher" json:"teacher"`
	Overview     string `db:"overview" json:"overview"`
	Goal         string `db:"goal" json:"goal"`
	EvaluateID   string `db:"evaluate_id"`
	Evaluate     []Evaluate
	Scehdule     []Scehdule
	Textbook     string `db:"textbook" json:"textbook"`
	ReferenceURL string `db:"reference_url" json:"reference_url"`
	Remarks      string `db:"remarks" json:"remarks"`
}

// フロントに返してやるモデル
type Evaluate struct {
	Method     string `db:"method" json:"method"`
	Comment    string `db:"comment" json:"comment"`
	Percentage string `db:"percentage" json:"percentage"`
}

type Scehdule struct {
	Session string `db:"session" json:"session"`
}
