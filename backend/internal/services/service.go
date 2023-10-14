package services

import (
	"fmt"
	"github.com/team-for-more-tech-5-0/opti-bank-backend.git/internal/database"
	"github.com/team-for-more-tech-5-0/opti-bank-backend.git/internal/models/bank"
	"log"
	"math"
	"sort"
)

func CalculateNearBanks(lat, lon float64, radius float64, service string) ([]bank.Bank, error) {
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
	for _, currentBank := range result {
		fmt.Printf("time: %d\n", currentBank.TotalTime)
	}
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

//func iterate(v interface{}) {
//	switch val := v.(type) {
//	case map[string]map[string]interface{}:
//		fmt.Println("map2:")
//		for k1, v1 := range val {
//			fmt.Printf("  %s:\n\n", k1)
//			for k2, v2 := range v1 {
//				fmt.Printf("    %s: %v\n----\n", k2, v2)
//			}
//		}
//	default:
//		fmt.Println("unknown type:", val)
//	}
//}

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
