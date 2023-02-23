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
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categorySerivice := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categorySerivice)
	router := app.NewRouter(categoryController)

	router.PanicHandler = excpetion.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
