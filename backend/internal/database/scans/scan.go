package scans

import (
	"database/sql/driver"
	"errors"
	"github.com/team-for-more-tech-5-0/opti-bank-backend.git/internal/models/service"
)

func (s service.Service) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Implement the Scanner interface
func (s *Service) Scan(value interface{}) error {
	if value == nil {
		*s = Service{}
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid JSONB data type")
	}
	return json.Unmarshal(b, s)
}
