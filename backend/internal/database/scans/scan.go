package scans

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/team-for-more-tech-5-0/opti-bank-backend.git/internal/models/service"
)

type MyService struct {
	service.Service
}

func (s MyService) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Implement the Scanner interface
func (s *MyService) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid JSONB data type")
	}
	return json.Unmarshal(b, s)
}
