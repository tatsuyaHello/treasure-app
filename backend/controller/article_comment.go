package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/httputil"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/service"
	"net/http"
	"strconv"
)

type ArticleComment struct {
	dbx *sqlx.DB
}

func NewArticleComment(dbx *sqlx.DB) *ArticleComment {
	return &ArticleComment{dbx: dbx}
}

func (ac *ArticleComment) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["article_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	createArticleComment := &model.CreateRequestArticleComment{}
	if err := json.NewDecoder(r.Body).Decode(&createArticleComment); err != nil {
		return http.StatusBadRequest, nil, err
	}

	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	newArticleComment := &model.ArticleComment{}
	newArticleComment.ID = aid
	newArticleComment.UserID = user.ID
	newArticleComment.ArticleID = aid

	articleCommentService := service.NewArticleCommentService(ac.dbx)
	createdId, err := articleCommentService.Create(newArticleComment)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newArticleComment.ID = createdId

	return http.StatusCreated, newArticleComment, nil
}
