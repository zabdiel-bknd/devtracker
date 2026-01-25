package main

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
	Priority string `json:"priority"`
	Status string `json:"status"` 			// Values: LOW, MEDIUM, HIGH
	ProjectID string `json:"project_id"`	// Values: TODO, DOING, DONE
	CreatedAt time.Time `json:"created_at"`
}