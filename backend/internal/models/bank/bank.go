package bank

type Bank struct {
	ID              int64                  `json:"bank_id"`
	SalePointName   string                 `json:"salepointname"`
	Address         string                 `json:"address"`
	Status          string                 `json:"status"`
	OpenHours       []Schedule             `json:"openhours"`
	Rko             string                 `json:"rko"`
	OfficeType      string                 `json:"officetype"`
	SalePointFormat string                 `json:"salepointformat"`
	Suoavailability string                 `json:"suoavailability"`
	Hasramp         string                 `json:"hasramp"`
	Latitude        float64                `json:"latitude"`
	Longitude       float64                `json:"longitude"`
	Metrostation    string                 `json:"metrostation"`
	Distance        string                 `json:"distance"`
	Kep             bool                   `json:"kep"`
	MyBranch        bool                   `json:"mybranch"`
	Service         map[string]interface{} `json:"service"`
}

//func NewBank(
//	id int64,
//	salePointName string,
//	address string,
//	status string,
//	schedule []Schedule,
//	rko string,
//	officeType string,
//	salePointFormat string,
//	suoavailability string,
//	hasramp string,
//	latitude float64,
//	longitude float64,
//	metrostation string,
//	distance string,
//	service service.Services,
//) *Bank {
//	return &Bank{
//		ID:              id,
//		SalePointName:   salePointName,
//		Address:         address,
//		Status:          status,
//		OpenHours:       schedule,
//		Rko:             rko,
//		OfficeType:      officeType,
//		SalePointFormat: salePointFormat,
//		Suoavailability: suoavailability,
//		Hasramp:         hasramp,
//		Latitude:        latitude,
//		Longitude:       longitude,
//		Metrostation:    metrostation,
//		Distance:        distance,
//		Service:         service,
//	}
//}
