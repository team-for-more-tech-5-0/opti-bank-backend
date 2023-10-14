package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/team-for-more-tech-5-0/opti-bank-backend.git/internal/database"
	"github.com/team-for-more-tech-5-0/opti-bank-backend.git/internal/services"
	"log"
	"net/http"
	"strconv"
)

func Transport() {
	router := gin.Default()
	router.GET("/getAllBanks", GetAllBanks)
	router.GET("/getAllAtms", GetAllAtms)
	router.GET("/getNearBanks", GetNearBanks)
	//router.GET("/getNearBanks", GetNearAtms)

	err := router.Run(":8088")
	if err != nil {
		panic(err)
	}
}

func GetAllBanks(context *gin.Context) {
	dbs, err := database.GetBanks()
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Cannot connect to dataBase"})
		return
	}
	context.IndentedJSON(http.StatusOK, dbs)
	//fmt.Println(dbs[0].Service)
	//fmt.Println("000000")
	//fmt.Println(dbs[0].Service["service"].(map[string]interface{})["servicesForBusinesses"].(map[string]interface{})["SMEservices"])

	return
}

func GetAllAtms(context *gin.Context) {
	dbs, err := database.GetAtms()
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Cannot connect to dataBase"})
		return
	}
	context.IndentedJSON(http.StatusOK, dbs)
	return
}

func GetNearBanks(context *gin.Context) {
	latitudeStr := context.Query("latitude")
	longitudeStr := context.Query("longitude")
	service := context.Query("service")

	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid latitude number to convert float"})
		log.Println(err)
		return
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid latitude number to convert float"})
		log.Println(err)
		return
	}
	dbs, err := services.CalculateNearBanks(latitude, longitude, 1, service)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot find near banks"})
		return
	}

	context.IndentedJSON(http.StatusOK, dbs)
	return
}

//func GetNearAtms(context *gin.Context) {
//
//}
