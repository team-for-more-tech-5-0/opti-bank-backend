package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/team-for-more-tech-5-0/opti-bank-backend.git/internal/database"
	"net/http"
)

func Transport() {
	router := gin.Default()
	router.GET("/getAllBanks", GetAllBanks)

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
	return
}
