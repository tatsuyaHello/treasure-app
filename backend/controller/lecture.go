package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/voyagegroup/treasure-app/httputil"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
	"github.com/voyagegroup/treasure-app/service"
)

type Lecture struct {
	dbx *sqlx.DB
}

type Review struct {
	dbx *sqlx.DB
}

func NewLecture(dbx *sqlx.DB) *Lecture {
	return &Lecture{dbx: dbx}
}

func NewReview(dbx *sqlx.DB) *Review {
	return &Review{dbx: dbx}
}

func (a *Lecture) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	lectures, err := repository.AllLecture(a.dbx)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, lectures, nil
}

func (a *Lecture) Search(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	// URLからクエリパラメータを取得する
	title := r.URL.Query().Get("title")
	english_title := r.URL.Query().Get("english_title")
	var unit int64
	var err error
	if r.URL.Query().Get("unit") != "" {
		unit, err = strconv.ParseInt(r.URL.Query().Get("unit"), 10, 64)
		if err != nil {
			return http.StatusBadRequest, nil, err
		}
	} else {
		// 適当な値
		unit = 100
	}

	semester := r.URL.Query().Get("semester")
	location := r.URL.Query().Get("location")
	lecture_style := r.URL.Query().Get("lecture_style")
	teacher := r.URL.Query().Get("teacher")
	// overviewは講義内容検索
	overview := r.URL.Query().Get("overview")

	// 何も指定しない場合は、BadRequestを返す
	if title == "" && english_title == "" && semester == "" && location == "" && lecture_style == "" && teacher == "" && overview == "" && unit == 100 {
		return http.StatusBadRequest, nil, err
	}

	lectures, err := repository.FindLecture(a.dbx, title, english_title, semester, location, lecture_style, teacher, overview, unit)
	if err != nil && err == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, lectures, nil
}

func (a *Review) CreateReview(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newReview := &model.Review{}
	if err := json.NewDecoder(r.Body).Decode(&newReview); err != nil {
		return http.StatusBadRequest, nil, err
	}

	// LectureIDをPATHから取得
	vars := mux.Vars(r)
	lectureID, ok := vars["lecture_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	newReview.LectureID = lectureID

	reviewService := service.NewReviewService(a.dbx)

	id, err := reviewService.Create(newReview)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	newReview.ID = id
	return http.StatusCreated, newReview, nil
}

func (a *Review) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["lecture_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	reviews, err := repository.GetReviews(a.dbx, id)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, reviews, nil
}
