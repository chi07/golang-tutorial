package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang-tutorial/go-clean-architecture/entities"
	"golang-tutorial/go-clean-architecture/packages/book"
	"golang-tutorial/go-clean-architecture/utils"
	"net/http"
	"strconv"
)

func FetchAll( s *book.Service) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
		books, err := s.FetchAll()
		if err != nil {
			utils.Respond(w, utils.Message(http.StatusInternalServerError, err.Error()))
			return
		}

		resp := utils.Message(http.StatusOK, "Get book info successfully!")
		resp["book"] = books
		utils.Respond(w, resp)
	})
}

func GetById( s *book.Service) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
		params := mux.Vars(r)
		bookId, err := strconv.ParseUint(string(params["id"]), 10, 64)
		if err != nil {
			utils.Respond(w, utils.Message(false, "ID is not valid"))
			return
		}

		b, err := s.FindById(uint(bookId))
		if  err != nil {
			utils.Respond(w, utils.Message(false, err.Error()))
			return
		}

		resp := utils.Message(http.StatusOK, "Get book info successfully!")
		resp["book"] = b
		utils.Respond(w, resp)
	})
}

func Store( s *book.Service) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
		request := entities.Book{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			logrus.Print(err)
			utils.Respond(w, utils.Message(http.StatusBadRequest, errors.New("Bad Request").Error()))
			return
		}

		// validate here

		// store
		if ok, err := s.Store(&request); err != nil {
			utils.Respond(w, utils.Message(ok, err.Error()))
			return
		}

		utils.Respond(w, utils.Message(http.StatusOK, string("Saved")))
		return
	})
}

func Update( s *book.Service) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
		params := mux.Vars(r)
		bookId, err := strconv.ParseUint(string(params["id"]), 10, 64)
		if err != nil {
			utils.Respond(w, utils.Message(false, "ID is not valid"))
			return
		}

		request := entities.Book{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			logrus.Print(err)
			utils.Respond(w, utils.Message(http.StatusBadRequest, errors.New("Bad Request").Error()))
			return
		}

		b, err := s.FindById(uint(bookId))
		if  err != nil {
			utils.Respond(w, utils.Message(http.StatusBadRequest, "Book not found!"))
			return
		}

		b.Name 		= request.Name
		b.Author 	= request.Author
		b.Year 		= request.Year

		if ok, err := s.Update(b); !ok || err != nil {
			utils.Respond(w, utils.Message(http.StatusBadRequest, err.Error()))
			return
		}

		utils.Respond(w, utils.Message(http.StatusOK, string("Saved")))
		return
	})
}

func Delete( s *book.Service) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
		params := mux.Vars(r)
		bookId, err := strconv.ParseUint(string(params["id"]), 10, 64)
		if err != nil {
			utils.Respond(w, utils.Message(http.StatusBadRequest, "ID is not valid"))
			return
		}

		if ok, err := s.Delete(uint(bookId)); !ok || err != nil {
			utils.Respond(w, utils.Message(http.StatusInternalServerError, err.Error()))
			return
		}

		utils.Respond(w, utils.Message(http.StatusOK, "Deleted"))
		return
	})
}

func MakeBookHandler(router *mux.Router, service *book.Service)  {
	router.Handle("/api/book", FetchAll(service)).Methods("GET")
	router.Handle("/api/book/{id}", GetById(service)).Methods("GET")
	router.Handle("/api/book/add", Store(service)).Methods("POST")
	router.Handle("/api/book/update/{id}", Update(service)).Methods("PUT", "PATCH")
	router.Handle("/api/book/delete/{id}", Delete(service)).Methods("DELETE")
}