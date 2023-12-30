package main

import (
	"fidibo-search/config"
	"fidibo-search/controller"
	_ "fidibo-search/docs"
	"fidibo-search/helper"
	"fidibo-search/model"
	"fidibo-search/repository"
	"fidibo-search/router"
	"fidibo-search/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// @title 	Book Service API
// @version	1.0
// @description A Book service API

// @host 	localhost:8888
// @BasePath /api
func main() {

	log.Info().Msg("Started Server!")

	cache := config.NewRedisClient()
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("books").AutoMigrate(&model.Books{})

	// Repository
	booksRepository := repository.NewBooksRepositoryImpl(db, cache)

	// Service
	booksService := service.NewBooksServiceImpl(booksRepository, validate)

	// Controller
	booksController := controller.NewBooksController(booksService)

	// Router
	routes := router.NewRouter(booksController)

	server := &http.Server{
		Addr:    config.APP_ADDRESS,
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
