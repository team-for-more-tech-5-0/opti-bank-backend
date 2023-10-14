package services

import (
	"fmt"
	"github.com/team-for-more-tech-5-0/opti-bank-backend.git/internal/database"
	"github.com/team-for-more-tech-5-0/opti-bank-backend.git/internal/models/bank"
	"log"
	"sort"
)

// Поиск ближайщих отделений
func CalculateNearBanks(lat, lon float64, radius float64, service string) ([]bank.Bank, error) {
	var result []bank.Bank
	radiusStep := 1.0

	for len(result) < 3 {
		newResult, err := FindNearBanks(lat, lon, radius, service)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		result = newResult
		radius += radiusStep
		radiusStep *= 1.5
	}
	return result, nil
}
func FindNearBanks(lat, lon float64, radius float64, service string) ([]bank.Bank, error) {
	var result []bank.Bank

	banks, err := database.GetBanks()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for _, currentBank := range banks {
		dist := distance(lat, lon, currentBank.Latitude, currentBank.Longitude)
		if dist <= radius && isServiceAvailable(service, currentBank.Service, &currentBank) {
			if currentBank.CurrentTypeService == "servicesForBusinesses" {
				currentBank.TotalTime = currentBank.TimeBusiness * currentBank.QueueBusiness
			} else {
				currentBank.TotalTime = currentBank.TimeIndividual * currentBank.QueueIndividual
			}
			result = append(result, currentBank)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].TotalTime < result[j].TotalTime
	})

	return result, nil
}

func isServiceAvailable(s string, service map[string]map[string]map[string]map[string]interface{}, currentBank *bank.Bank) bool {
	for _, value := range service {
		for serviceType, v1 := range value {
			fmt.Printf("serviceType:%s\n\n", serviceType)
			for key, v2 := range v1 {
				fmt.Printf("%v\n----\n", v2["serviceActivity"])
				if key == s && v2["serviceActivity"] == "AVAILABLE" {
					currentBank.CurrentTypeService = serviceType
					return true
				}
			}
		}
	}

	return false
}
