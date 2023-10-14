package transport

import (
	"github.com/Spawn4real/hackthon_more.tech_5.0.git/internal/database"
	"github.com/gin-gonic/gin"
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
	}
	context.IndentedJSON(http.StatusOK, dbs)
	return
}
