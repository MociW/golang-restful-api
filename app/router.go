package app

import (
	"Mociw/golang-api/controller"
	"Mociw/golang-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(CategoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", CategoryController.FindAll)
	router.GET("/api/categories/:categoryId", CategoryController.FindById)
	router.POST("/api/categories", CategoryController.Create)
	router.PUT("/api/categories/:categoryId", CategoryController.Update)
	router.DELETE("/api/categories/:categoryId", CategoryController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}
