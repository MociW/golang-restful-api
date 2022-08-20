package test

import (
	"database/sql"
	"net/http"
	"time"

	"Mociw/golang-api/app"
	"Mociw/golang-api/controller"
	"Mociw/golang-api/helper"
	"Mociw/golang-api/middleware"
	"Mociw/golang-api/repository"
	"Mociw/golang-api/service"

	"github.com/go-playground/validator"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/data_user?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}

func setupTouter() http.Handler {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	CategoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(CategoryController)

	return middleware.NewAuthMidlleware(router)
}

func Test
