package models

type Task struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Completed   bool   `json:"completed"`
}
