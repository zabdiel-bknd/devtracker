package models

import "time"

// Project represents a high-level project in the database.
type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
}

// Task represents a specific unit of work within a project.
type Task struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Priority string `json:"priority"`		// Values: LOW, MEDIUM, HIGH
	Status string `json:"status"` 			// Values: TODO, DOING, DONE
	ProjectID int `json:"project_id"`	
	CreatedAt time.Time `json:"created_at"`
}

func (task *Task) IsValid() bool {
	if task.Title == "" {
		return false
	}

	validPriorities := map[string]bool{"LOW": true, "MEDIUM": true, "HIGH": true}
	if !validPriorities[task.Priority]{
		return false
	}

	validStatuses := map[string]bool{"TODO": true, "DOING":true, "DONE": true}
	if !validStatuses[task.Status]{
		return false
	}

	return true
}