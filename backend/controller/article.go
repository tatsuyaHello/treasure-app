package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/voyagegroup/treasure-app/httputil"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
	"github.com/voyagegroup/treasure-app/service"
)

type Article struct {
	dbx *sqlx.DB
}

// type ArticleTags struct {
// 	dbx *sqlx.DB
// }

type Comment struct {
	dbx *sqlx.DB
}

func NewArticle(dbx *sqlx.DB) *Article {
	return &Article{dbx: dbx}
}

// func NewArticleTags(dbx *sqlx.DB) *ArticleTags {
// 	return &ArticleTags{dbx: dbx}
// }

func NewComment(dbx *sqlx.DB) *Comment {
	return &Comment{dbx: dbx}
}

func (a *Article) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	articles, err := repository.AllArticle(a.dbx)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, articles, nil
}

func (a *Article) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	article, err := repository.FindArticle(a.dbx, aid)
	if err != nil && err == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, article, nil
}

func (a *Comment) CreateComment(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newComment := &model.Comment{}
	if err := json.NewDecoder(r.Body).Decode(&newComment); err != nil {
		return http.StatusBadRequest, nil, err
	}

	// ArticleIDをPATHから取得
	vars := mux.Vars(r)
	articleID, ok := vars["article_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	newComment.ArticleID, _ = strconv.ParseInt(articleID, 10, 64)

	commentService := service.NewCommentService(a.dbx)

	// UserIDの取得
	user := &model.User{}
	user, err := httputil.GetUserFromContext(r.Context())
	newComment.UserID = user.ID

	id, err := commentService.Create(newComment)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	newComment.ID = id
	return http.StatusCreated, newComment, nil
}

func (a *Article) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	requestArticle := &model.CreateArticleRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestArticle); err != nil {
		return http.StatusBadRequest, nil, err
	}
	articleService := service.NewArticleService(a.dbx)

	newArticle := &model.Article{}
	newArticle.Body = requestArticle.Body
	newArticle.Title = requestArticle.Title

	user := &model.User{}
	user, err := httputil.GetUserFromContext(r.Context())
	newArticle.UserID = &user.ID
	id, err := articleService.Create(newArticle, requestArticle.TagIDs)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newArticle.ID = id

	responseArticle := &model.CreateArticleResponse{}
	responseArticle.ID = newArticle.ID
	responseArticle.Title = newArticle.Title
	responseArticle.Body = newArticle.Body
	responseArticle.UserID = newArticle.UserID
	responseArticle.Tags = requestArticle.TagIDs

	return http.StatusCreated, responseArticle, nil
}

func (a *Article) Update(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	reqArticle := &model.Article{}
	if err := json.NewDecoder(r.Body).Decode(&reqArticle); err != nil {
		return http.StatusBadRequest, nil, err
	}

	articleService := service.NewArticleService(a.dbx)
	err = articleService.Update(aid, reqArticle)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}

func (a *Article) Destroy(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	articleService := service.NewArticleService(a.dbx)
	err = articleService.Destroy(aid)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}
