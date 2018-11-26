package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"golang-tutorial/go-clean-architecture/packages/book"
	delivery "golang-tutorial/go-clean-architecture/packages/book/handler/http"
	"golang-tutorial/go-clean-architecture/database"
	"log"
)

func main()  {
	router := mux.NewRouter()
	mysqlDb := database.GetMysqlConnectionPool()

	// user handler
	bookRepo := book.NewMysqlBookRepository(mysqlDb)
	bookService := book.NewService(bookRepo)
	delivery.MakeBookHandler(router, bookService)

	log.Fatal(http.ListenAndServe(":8090", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
