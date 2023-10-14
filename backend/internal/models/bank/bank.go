package bank

type Bank struct {
	ID              int64      `json:"bank_id"`
	SalePointName   string     `json:"salepointname"`
	Address         string     `json:"address"`
	Status          string     `json:"status"`
	OpenHours       []Schedule `json:"openhours"`
	Rko             string     `json:"rko"`
	OfficeType      string     `json:"officetype"`
	SalePointFormat string     `json:"salepointformat"`
	Suoavailability string     `json:"suoavailability"`
	Hasramp         string     `json:"hasramp"`
	Latitude        float64    `json:"latitude"`
	Longitude       float64    `json:"longitude"`
	Metrostation    string     `json:"metrostation"`
	Distance        string     `json:"distance"`
	Kep             bool       `json:"kep"`
	MyBranch        bool       `json:"mybranch"`
	QueueIndividual int64      `json:"queueIndividual"`
	QueueBusiness   int64      `json:"queueBusiness"`
	TimeIndividual  int64      `json:"timeIndividual"`
	TimeBusiness    int64      `json:"timeBusiness"`
	//для конвертации из типа jsonb-jsonb,jsonb....
	Service            map[string]map[string]map[string]map[string]interface{} `json:"service"`
	TotalTime          int64                                                   `json:"totalTime"`
	CurrentTypeService string                                                  `json:"currentTypeService"`
}
