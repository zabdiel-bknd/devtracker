package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/zabdiel-bknd/devtracker/internal/database"
	"github.com/zabdiel-bknd/devtracker/internal/models"
)

type TaskHandler struct{
	store *database.Service
}

func NewTaskHandler(store *database.Service) *TaskHandler{
	return &TaskHandler{store: store}
}

func (handler *TaskHandler) Create(w http.ResponseWriter, r *http.Request){
	projectIDStr := r.PathValue("id")
	projectID, err := strconv.Atoi(projectIDStr)

	if err != nil{
		http.Error(w, "Invalid Project ID", http.StatusBadRequest)
	}

	var t models.Task

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil{
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
	}

	t.ProjectID = projectID

	if err := handler.store.CreateTask(&t); err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)

}

func (handler *TaskHandler) List(w http.ResponseWriter, r *http.Request){
	projectIDStr := r.PathValue("id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		http.Error(w, "Invalid Project ID", http.StatusBadRequest)
		return
	}

	tasks, err := handler.store.GetTasksByProject(projectID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if tasks == nil{
		tasks = []models.Task{} //return empty array
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}