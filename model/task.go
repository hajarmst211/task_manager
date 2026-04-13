package model

import (
	"time"
)

type Task struct {
	ID       int       `json:"id"`       // 2. Use backticks, not single quotes
	Title    string    `json:"title"`    
	Deadline time.Time `json:"deadline"` 
	Details  string    `json:"details"`  
	Done     bool      `json:"done"`     
}




