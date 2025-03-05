package storage

import (
	"sync"

	"github.com/nbursa/whistlechain-backend/models"
)

// In-memory storage for reports
type ReportStore struct {
	mu      sync.Mutex
	reports map[string]models.Report
}

// NewReportStore initializes an empty report store
func NewReportStore() *ReportStore {
	return &ReportStore{
		reports: make(map[string]models.Report),
	}
}

// AddReport saves a report in memory
func (s *ReportStore) AddReport(report models.Report) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.reports[report.ID] = report
}

// GetAllReports returns all stored reports
func (s *ReportStore) GetAllReports() []models.Report {
	s.mu.Lock()
	defer s.mu.Unlock()
	var result []models.Report
	for _, r := range s.reports {
		result = append(result, r)
	}
	return result
}

// GetReportByID returns a report by its ID
func (s *ReportStore) GetReportByID(id string) (models.Report, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	report, exists := s.reports[id]
	return report, exists
}
