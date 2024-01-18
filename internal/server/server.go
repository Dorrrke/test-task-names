package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Dorrrke/test-task-names/internal/domain/models"
	"github.com/Dorrrke/test-task-names/internal/logger"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Server struct {
	HttpServer        *http.Server
	storage           Storage
	enrichmentService Enrichment
	// Клиенты для работы обогащения
}

type Storage interface {
	GetNameById(ctx context.Context, id int) (models.NameData, error)
	GetAllNames(ctx context.Context) ([]models.NameData, error)
	GetNameBySurname(ctx context.Context, surname string) (models.NameData, error)
	GetNamesByPatronymic(ctx context.Context, patronymic string) (models.NameData, error)
	GetNamesByAge(ctx context.Context, age int) (models.NameData, error)
	GetNamesByGender(ctx context.Context, gender string) (models.NameData, error)
	GetNamesByNational(ctx context.Context, national string) (models.NameData, error)
	SaveName(ctx context.Context, name models.NameData) error
	DeleteName(ctx context.Context, id int) error
	UpdateName(ctx context.Context, name models.NameData, id int) error
}

type Enrichment interface {
	EnrichName(name models.NameData) (models.NameData, error)
}

func New(serv *http.Server, storage Storage, enrichService Enrichment) *Server {

	return &Server{
		HttpServer:        serv,
		storage:           storage,
		enrichmentService: enrichService,
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
			r.Get("/", s.ByIdHandler)
			r.Get("/update", s.UpdateNameHandler)
		})
		r.Route("/users/filter", func(r chi.Router) {
			r.Get("/surmane", s.BySurNameHandler)
			r.Get("/patronymic", s.ByPatronymicHandler)
			r.Get("/age", s.ByAgeHandler)
			r.Get("/gender", s.ByGenderHandler)
			r.Get("/national", s.ByNationalHandler)
		})
	})

	s.HttpServer.Handler = r
}

func (s *Server) ByIdHandler(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	if id != "" {
		nameID, err := strconv.Atoi(id)
		if err != nil {
			logger.Log.Error("Error convert id to int", zap.Error(err))
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*160))
		defer cancel()
		name, err := s.storage.GetNameById(ctx, nameID)
		if err != nil {
			//TODO: Обработтка несуществующего пользователя
			logger.Log.Error("Error convert id to int", zap.Error(err))
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusCreated)
		enc := json.NewEncoder(res)
		if err := enc.Encode(name); err != nil {
			logger.Log.Debug("error encoding responce", zap.Error(err))
			http.Error(res, "Internal Error", http.StatusInternalServerError)
			return
		}
		return
	}
	logger.Log.Error("id is empty")
	http.Error(res, "Не корректный запрос", http.StatusBadRequest)
}

func (s *Server) BySurNameHandler(res http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	var nameModel models.NameData

	if err := dec.Decode(&nameModel); err != nil {
		logger.Log.Error("Error decode body", zap.Error(err))
		http.Error(res, "Не корректный запрос", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*160))
	defer cancel()
	name, err := s.storage.GetNameBySurname(ctx, nameModel.Surname)
	if err != nil {
		//TODO: Обработтка несуществующего пользователя
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	enc := json.NewEncoder(res)
	if err := enc.Encode(name); err != nil {
		logger.Log.Debug("error encoding responce", zap.Error(err))
		http.Error(res, "Internal Error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) AllNamesHandler(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*160))
	defer cancel()
	names, err := s.storage.GetAllNames(ctx)
	if err != nil {
		//TODO: Обработтка несуществующего пользователя
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if len(names) == 0 {
		http.Error(res, "Нет сохраненных адресов", http.StatusNoContent)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(res)
	if err := enc.Encode(names); err != nil {
		logger.Log.Debug("error encoding responce", zap.Error(err))
		http.Error(res, "Не корректный запрос", http.StatusInternalServerError)
		return
	}
}

func (s *Server) ByPatronymicHandler(res http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	var nameModel models.NameData

	if err := dec.Decode(&nameModel); err != nil {
		logger.Log.Error("Error decode body", zap.Error(err))
		http.Error(res, "Не корректный запрос", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*160))
	defer cancel()
	name, err := s.storage.GetNamesByPatronymic(ctx, nameModel.Patronymic)
	if err != nil {
		//TODO: Обработтка несуществующего пользователя
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	enc := json.NewEncoder(res)
	if err := enc.Encode(name); err != nil {
		logger.Log.Debug("error encoding responce", zap.Error(err))
		http.Error(res, "Internal Error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) ByAgeHandler(res http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	var nameModel models.NameData

	if err := dec.Decode(&nameModel); err != nil {
		logger.Log.Error("Error decode body", zap.Error(err))
		http.Error(res, "Не корректный запрос", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*160))
	defer cancel()
	name, err := s.storage.GetNamesByAge(ctx, nameModel.Age)
	if err != nil {
		//TODO: Обработтка несуществующего пользователя
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	enc := json.NewEncoder(res)
	if err := enc.Encode(name); err != nil {
		logger.Log.Debug("error encoding responce", zap.Error(err))
		http.Error(res, "Internal Error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) ByGenderHandler(res http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	var nameModel models.NameData

	if err := dec.Decode(&nameModel); err != nil {
		logger.Log.Error("Error decode body", zap.Error(err))
		http.Error(res, "Не корректный запрос", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*160))
	defer cancel()
	name, err := s.storage.GetNamesByGender(ctx, nameModel.Gender)
	if err != nil {
		//TODO: Обработтка несуществующего пользователя
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	enc := json.NewEncoder(res)
	if err := enc.Encode(name); err != nil {
		logger.Log.Debug("error encoding responce", zap.Error(err))
		http.Error(res, "Internal Error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) ByNationalHandler(res http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	var nameModel models.NameData

	if err := dec.Decode(&nameModel); err != nil {
		logger.Log.Error("Error decode body", zap.Error(err))
		http.Error(res, "Не корректный запрос", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*160))
	defer cancel()
	name, err := s.storage.GetNamesByNational(ctx, nameModel.National)
	if err != nil {
		//TODO: Обработтка несуществующего пользователя
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	enc := json.NewEncoder(res)
	if err := enc.Encode(name); err != nil {
		logger.Log.Debug("error encoding responce", zap.Error(err))
		http.Error(res, "Internal Error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) SaveNameHandler(res http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	var nameModel models.NameData

	if err := dec.Decode(&nameModel); err != nil {
		logger.Log.Error("Error decode body", zap.Error(err))
		http.Error(res, "Не корректный запрос", http.StatusBadRequest)
		return
	}

	result, err := s.enrichmentService.EnrichName(nameModel)
	if err != nil {
		logger.Log.Error("Error enric name", zap.Error(err))
		http.Error(res, "Internal Error", http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*160))
	defer cancel()
	err = s.storage.SaveName(ctx, result)
	if err != nil {
		logger.Log.Debug("error save name", zap.Error(err))
		http.Error(res, "Internal Error", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	enc := json.NewEncoder(res)
	if err := enc.Encode(result); err != nil {
		logger.Log.Debug("error encoding responce", zap.Error(err))
		http.Error(res, "Internal Error", http.StatusInternalServerError)
		return
	}

}

func (s *Server) DeleteNameHandler(res http.ResponseWriter, req *http.Request) {

}

func (s *Server) UpdateNameHandler(res http.ResponseWriter, req *http.Request) {

}
