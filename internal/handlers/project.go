package handlers

import (
	"encoding/json"
	"net/http"
	"log/slog"
	"strconv"
	"database/sql"

	"github.com/zabdiel-bknd/devtracker/internal/database"
	"github.com/zabdiel-bknd/devtracker/internal/models"

)

// ProjectHandler groups all endpoints related to PROJECTS.
//set 'store inside the struct, so all methods (Create, Get, etc.) have acces to DB

type ProjectHandler struct {
	store *database.Service
}

// Constructor that receives the DB dependency
func NewProjectHandler(store *database.Service) *ProjectHandler {
	return &ProjectHandler{
		store: store,
	}
}


func (handler *ProjectHandler) Create(w http.ResponseWriter, r *http.Request){
	var project models.Project

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		slog.Warn("Failed to decode JSON", "error", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := handler.store.CreateProject(&project); err != nil{
		slog.Error("Database Insert Failed", 
            "error", err, 
            "project_name", project.Name, 
        )
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	slog.Info("Project created successfully", "id", project.ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)

}


func (handler *ProjectHandler) GetById(w http.ResponseWriter, r *http.Request){
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
	}

	project, err := handler.store.GetProject(id)
	if err != nil {
		if err == sql.ErrNoRows{
			http.Error(w, "Project Not found", http.StatusNotFound)
		}else{
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}
