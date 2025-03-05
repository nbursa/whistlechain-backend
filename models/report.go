package models

import (
	"time"

	"github.com/google/uuid"
)

// Report represents an anonymous report
type Report struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
}

// NewReport creates a new report with a unique ID
func NewReport(description string) Report {
	return Report{
		ID:          uuid.New().String(),
		Description: description,
		Timestamp:   time.Now(),
	}
}
