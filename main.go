package main

import (
	"github.com/Ahmad-mufied/learn-golang-restful-api/app"
	"github.com/Ahmad-mufied/learn-golang-restful-api/controller"
	"github.com/Ahmad-mufied/learn-golang-restful-api/excpetion"
	"github.com/Ahmad-mufied/learn-golang-restful-api/helper"
	"github.com/Ahmad-mufied/learn-golang-restful-api/middleware"
	"github.com/Ahmad-mufied/learn-golang-restful-api/repository"
	"github.com/Ahmad-mufied/learn-golang-restful-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categorySerivice := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categorySerivice)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = excpetion.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
