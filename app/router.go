package app

import (
	"github.com/Ahmad-mufied/learn-golang-restful-api/controller"
	"github.com/Ahmad-mufied/learn-golang-restful-api/excpetion"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = excpetion.ErrorHandler

	return router
}
