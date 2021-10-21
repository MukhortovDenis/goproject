package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Функция для работы с go-chi роутером
func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	return router
}

// func NewHandler() http.Handler {
// 	handler := chi.NewRouter()
// 	return handler
// }
