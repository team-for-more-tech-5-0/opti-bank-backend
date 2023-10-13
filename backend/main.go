// main.go

package main

import (
    "net/http"

    "github.com/gin-gonic/gin" // Используем фреймворк Gin для создания веб-сервиса
)

func main() {
    router := gin.Default()

    // Здесь вы можете определить маршруты и обработчики запросов

    // Пример маршрута
    router.GET("/api/banks", func(c *gin.Context) {
        // Здесь можно обработать запрос и вернуть данные о банках
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, World!",
        })
    })

    // Запускаем сервер на порту 8080
    router.Run(":8080")
}
