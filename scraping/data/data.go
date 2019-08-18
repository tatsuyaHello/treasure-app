package data

// DBに保存する際に使用するモデル
// Lecture is
type Lecture struct {
	Year         string
	LectureID    string
	Title        string
	SubTitle     string
	EnglishTitle string
	Unit         int
	Semester     string
	Location     string
	LectureStyle string
	Teacher      string
	Overview     string
	Goal         string
	EvaluateID   string
	Textbook     string
	ReferenceURL string
	Remarks      string
}

type Scehdule struct {
	LectureID string
	Session   string
}

// Evaluate is
type Evaluate struct {
	ID         string
	Method     string
	Comment    string
	Percentage string
}
