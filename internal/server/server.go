package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	HttpServer *http.Server
	storage    Storage
	// Клиенты для работы обогащения
}

type Storage interface {
	GetNameById()
	GetAllNames()
	GetNameBySurname()
	GetNamesByPatronymic()
	GetNamesByAge()
	GetNamesBySex()
	GetNamesByNational()
	SaveName()
	DeleteName()
	UpdateName()
}

func New(serv *http.Server, storage Storage) *Server {

	return &Server{
		HttpServer: serv,
		storage:    storage,
	}
}

func (s *Server) RegisterServer() {
	r := chi.NewRouter()
	// TODO: Доделать роутер
	r.Route("/", func(r chi.Router) {
		r.Post("/new", s.SaveNameHandler)
		r.Delete("/delete", s.DeleteNameHandler)
		r.Get("/users", s.AllNamesHandler)
		r.Route("/user/{id}", func(r chi.Router) {
			r.Get("/update", s.UpdateNameHandler)
		})
		r.Route("/users/filter", func(r chi.Router) {
			r.Get("/surmane", s.BySurNameHandler)
			r.Get("/patronymic", s.ByPatronymicHandler)
			r.Get("/age", s.ByAgeHandler)
			r.Get("/sex", s.BySexHandler)
			r.Get("/national", s.ByNationalHandler)
		})
	})

	s.HttpServer.Handler = r
}

func (s *Server) ByIdHandler(res http.ResponseWriter, req *http.Request) {

}

func (s *Server) BySurNameHandler(res http.ResponseWriter, req *http.Request) {

}

func (s *Server) AllNamesHandler(res http.ResponseWriter, req *http.Request) {

}

func (s *Server) ByPatronymicHandler(res http.ResponseWriter, req *http.Request) {

}

func (s *Server) ByAgeHandler(res http.ResponseWriter, req *http.Request) {

}

func (s *Server) BySexHandler(res http.ResponseWriter, req *http.Request) {

}

func (s *Server) ByNationalHandler(res http.ResponseWriter, req *http.Request) {

}

func (s *Server) SaveNameHandler(res http.ResponseWriter, req *http.Request) {

}

func (s *Server) DeleteNameHandler(res http.ResponseWriter, req *http.Request) {

}

func (s *Server) UpdateNameHandler(res http.ResponseWriter, req *http.Request) {

}
