package models

import "github.com/ServiceWeaver/weaver"

type Task struct {
	weaver.AutoMarshal
	ID          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
