package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/oxxi/jobsity/models"
	"github.com/oxxi/jobsity/services"
	"github.com/oxxi/jobsity/utils"
)

type ITaskController interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
	GetAllTask(w http.ResponseWriter, r *http.Request)
	GetTaskById(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
}

type taskController struct {
	service services.ITaskService
}

// validate
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// CreateTask implements ITaskController.
func (t *taskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	json.NewDecoder(r.Body).Decode(&task)
	// validate struct
	erro := validate.Struct(task)
	if erro != nil {
		for _, err := range erro.(validator.ValidationErrors) {
			message := "Field " + err.Field() + " is " + err.Tag()
			utils.RespondWithError(w, http.StatusBadRequest, message)
			return
		}
	}

	newTask, err := t.service.Save(r.Context(), task)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "error while creating task")
		return
	}
	result, err := json.Marshal(newTask)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to encode task")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

// DeleteTask implements ITaskController.
func (t *taskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "parameter Id cannot be parsed, please make sure the ID is valid")
		return
	}

	if err := t.service.Delete(r.Context(), id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "parameter Id cannot be parsed, please make sure the ID is valid")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// GetAllTask implements ITaskController.
func (t *taskController) GetAllTask(w http.ResponseWriter, r *http.Request) {

	tasks, err := t.service.GetAllTask(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	result, err := json.Marshal(tasks)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "request cannot be parsed, please make sure the request is valid")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

// GetTaskById implements ITaskController.
func (t *taskController) GetTaskById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "parameter Id cannot be parsed, please make sure the ID is valid")
		return
	}

	task, er := t.service.GetById(ctx, id)
	if er != nil {

		utils.RespondWithError(w, http.StatusInternalServerError, er.Error())
		return
	}

	result, err := json.Marshal(task)

	if err != nil {

		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to encode task")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

// UpdateTask implements ITaskController.
func (t *taskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	ctx := r.Context()

	//binding struct
	json.NewDecoder(r.Body).Decode(&task)
	//validate struct
	errStruct := validate.Struct(task)
	if errStruct != nil {
		for _, err := range errStruct.(validator.ValidationErrors) {
			message := "Field " + err.Field() + " is " + err.Tag()
			utils.RespondWithError(w, http.StatusBadRequest, message)
			return
		}
	}
	//validate get and path value
	id, errVal := strconv.Atoi(r.PathValue("id"))
	if errVal != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "parameter Id cannot be parsed, please make sure the ID is valid")
		return
	}

	// update new values
	newModel, err := t.service.Update(ctx, id, task)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	result, err := json.Marshal(newModel)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to encode task")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func NewTaskController(s services.ITaskService) ITaskController {
	return &taskController{
		service: s,
	}
}
