package main

import (
	"net/http"

	"Mociw/golang-api/app"
	"Mociw/golang-api/controller"
	"Mociw/golang-api/helper"
	"Mociw/golang-api/middleware"
	"Mociw/golang-api/repository"
	"Mociw/golang-api/service"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	CategoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(CategoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMidlleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
