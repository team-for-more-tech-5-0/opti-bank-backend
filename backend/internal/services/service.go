package services

import (
	"errors"
	"math"
)

func GetClientType(service string) (string, error) {
	if service == "SMEservices" || service == "businessFinancing" || service == "corporateAccounts" || service == "corporateCreditCards" || service == "transactionAndPaymentServices" || service == "internationalOperationsServices" || service == "liquidityAndFinancialRiskManagement" {
		return "servicesForBusinesses", nil
	} else if service == "mortgageLoans" || service == "loansAndCredits" || service == "depositsAndSavings" || service == "investmentServices" || service == "bankAccountsAndCards" || service == "onlineBankingAndMobileApp" {
		return "servicesForIndividuals", nil
	} else {
		return "", errors.New("unknown client type")
	}
}

// Поиск по широте и долготе в радиусе
func distance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Earth radius in km
	dLat := deg2rad(lat2 - lat1)
	dLon := deg2rad(lon2 - lon1)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(deg2rad(lat1))*math.Cos(deg2rad(lat2))*math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func deg2rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}
