package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"sync"

	"github.com/zabdiel-bknd/devtracker/internal/database"
)

type DashboardHanlder struct{
	store *database.Service
}

func NewDashboardHandler(store *database.Service) *DashboardHanlder{
	return &DashboardHanlder{store: store}
}

type DashboardResponse struct {
	TotalProjects int `json:"total_projects"`
	TotalTasks    int `json:"total_tasks"`
}

func (handler *DashboardHanlder) GetStats(w http.ResponseWriter, r *http.Request){
	var (
			projectsCount int
			tasksCount    int
			errProjects   error
			errTasks      error
		)
	
	// Initialize wait group for 2 tasks
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		defer wg.Done()
		slog.Info("Fetching projects count ...")
		projectsCount, errProjects = handler.store.GetProjectsCount()
	}()

	go func(){
		defer wg.Done()
		slog.Info("Fetching tasks count ...")
		tasksCount, errTasks = handler.store.GetTasksCount()
	}()

	wg.Wait()

	if errProjects != nil || errTasks != nil {
        slog.Error("Dashboard error", "proj_err", errProjects, "task_err", errTasks)
		http.Error(w, "Error calculating stats", http.StatusInternalServerError)
		return
	}

	resp := DashboardResponse{
		TotalProjects: projectsCount,
		TotalTasks:    tasksCount,
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}