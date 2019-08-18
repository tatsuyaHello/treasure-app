package server

import (
	"fmt"

	"github.com/voyagegroup/treasure-app/sample"

	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/justinas/alice"

	"firebase.google.com/go/auth"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
	"github.com/voyagegroup/treasure-app/controller"
	db2 "github.com/voyagegroup/treasure-app/db"
	"github.com/voyagegroup/treasure-app/firebase"
	"github.com/voyagegroup/treasure-app/middleware"
)

type Server struct {
	dbx        *sqlx.DB
	router     *mux.Router
	authClient *auth.Client
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(datasource string) {
	authClient, err := firebase.InitAuthClient()
	if err != nil {
		log.Fatalf("failed init auth client. %s", err)
	}
	s.authClient = authClient

	db := db2.NewDB(datasource)
	dbx, err := db.Open()
	if err != nil {
		log.Fatalf("failed db init. %s", err)
	}
	s.dbx = dbx
	s.router = s.Route()
}

func (s *Server) Run(addr string) {
	log.Printf("Listening on port %s", addr)
	err := http.ListenAndServe(
		fmt.Sprintf(":%s", addr),
		handlers.CombinedLoggingHandler(os.Stdout, s.router),
	)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route() *mux.Router {
	authMiddleware := middleware.NewAuthMiddleware(s.authClient, s.dbx)
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Content-Type"},
	})

	commonChain := alice.New(
		middleware.RecoverMiddleware,
		corsMiddleware.Handler,
	)

	authChain := commonChain.Append(
		authMiddleware.Handler,
	)

	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/public").Handler(commonChain.Then(sample.NewPublicHandler()))
	r.Methods(http.MethodGet).Path("/private").Handler(authChain.Then(sample.NewPrivateHandler(s.dbx)))

	articleController := controller.NewArticle(s.dbx)
	r.Methods(http.MethodPost).Path("/articles").Handler(authChain.Then(AppHandler{articleController.Create}))
	r.Methods(http.MethodPut).Path("/articles/{id}").Handler(authChain.Then(AppHandler{articleController.Update}))
	r.Methods(http.MethodDelete).Path("/articles/{id}").Handler(authChain.Then(AppHandler{articleController.Destroy}))
	r.Methods(http.MethodGet).Path("/articles").Handler(commonChain.Then(AppHandler{articleController.Index}))
	r.Methods(http.MethodGet).Path("/articles/{id}").Handler(commonChain.Then(AppHandler{articleController.Show}))

	lectureController := controller.NewLecture(s.dbx)
	r.Methods(http.MethodGet).Path("/lectures").Handler(commonChain.Then(AppHandler{lectureController.Index}))
	r.Methods(http.MethodGet).Path("/lecture").Handler(commonChain.Then(AppHandler{lectureController.Search}))

	reviewController := controller.NewReview(s.dbx)
	// 一旦、講義に対して誰でもレビューをすることができる状態にする
	r.Methods(http.MethodPost).Path("/lectures/{lecture_id}/reviews").Handler(commonChain.Then(AppHandler{reviewController.CreateReview}))

	r.Methods(http.MethodGet).Path("/reviews/{lecture_id}").Handler(commonChain.Then(AppHandler{reviewController.Index}))

	commentController := controller.NewComment(s.dbx)
	r.Methods(http.MethodPost).Path("/articles/{article_id}/comments").Handler(authChain.Then(AppHandler{commentController.CreateComment}))

	r.PathPrefix("").Handler(commonChain.Then(http.StripPrefix("/img", http.FileServer(http.Dir("./img")))))
	return r
}
