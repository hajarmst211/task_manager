package model

import (
	"time"
)

type Task struct {
	ID       int       `json:"id"`       
	Title    string    `json:"title"`  
	CreatedAt time.Time `json:"createdAt"`
	Deadline time.Time `json:"deadline"` 
	Details  string    `json:"details"`  
	Done     bool      `json:"done"`     
}




