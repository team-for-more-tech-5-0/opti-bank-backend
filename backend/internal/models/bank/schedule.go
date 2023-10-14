package bank

import "encoding/json"

type Schedule struct {
	Days  string `json:"days"`
	Hours string `json:"hours"`
}

func (s *Schedule) Scan(value interface{}) error {
	if value == nil {
		*s = Schedule{}
		return nil
	}
	var str string
	err := json.Unmarshal(value.([]byte), &str)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(str), s)
}
