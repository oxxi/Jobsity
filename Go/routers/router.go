package routers

import (
	"net/http"

	"github.com/oxxi/jobsity/controllers"
	"github.com/oxxi/jobsity/repositories"
	"github.com/oxxi/jobsity/services"
	"gorm.io/gorm"
)

var RegisterRouter = func(router *http.ServeMux, Db *gorm.DB) {
	repo := repositories.NewTaskRepository(Db)
	service := services.NewTaskService(repo)
	handler := controllers.NewTaskController(service)

	router.HandleFunc("POST /task", handler.CreateTask)
	router.HandleFunc("GET /tasks/", handler.GetAllTask)
	router.HandleFunc("GET /task/{id}", handler.GetTaskById)
	router.HandleFunc("PUT /task/{id}", handler.UpdateTask)
	router.HandleFunc("DELETE /task/{id}", handler.DeleteTask)

}
