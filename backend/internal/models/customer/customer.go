package customer

type Customer struct {
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	Preferences []string `json:"customer_preferences"`
}

//func NewCustomer(Latitude float64, Longitude float64) *Customer {
//
//}
