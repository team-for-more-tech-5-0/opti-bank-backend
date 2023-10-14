package atm

type Atm struct {
	ID        int64                  `json:"atm_id"`
	Address   string                 `json:"address"`
	Latitude  float64                `json:"latitude"`
	Longitude float64                `json:"longitude"`
	IsAllDay  bool                   `json:"allday"`
	Services  map[string]interface{} `json:"services"`
}

//func NewAtm(id int64, address string, latitude float64, longitude float64, isAllDay bool, services Services) *Atm {
//	return &Atm{
//		ID:        id,
//		Address:   address,
//		Latitude:  latitude,
//		Longitude: longitude,
//		IsAllDay:  isAllDay,
//		Services:  services,
//	}
//}
