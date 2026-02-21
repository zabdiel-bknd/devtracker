package models

import (
	"testing"
)
func TestTask_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		task     Task
		expected bool
	}{
		{
			name: "Valid Task",
			task: Task{Title: "Fix bug", Priority: "HIGH", Status: "TODO"},
			expected: true,
		},
		{
			name: "Missing Title",
			task: Task{Title: "", Priority: "MEDIUM", Status: "DONE"},
			expected: false,
		},
		{
			name: "Invalid Priority",
			task: Task{Title: "Write docs", Priority: "URGENT", Status: "TODO"},
			expected: false,
		},
		{
			name: "Invalid Status",
			task: Task{Title: "Deploy", Priority: "LOW", Status: "FINISHED"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			
			result := tt.task.IsValid()

			if result != tt.expected {
				t.Errorf("IsValid() = %v, want %v", result, tt.expected)
			}
		})
	}
}