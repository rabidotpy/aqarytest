package main

import (
	"aqarytest/database"
	"aqarytest/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    envErr := godotenv.Load()
    if envErr != nil {
      log.Fatal("Error loading .env file")
    }
    var databaseURL string = os.Getenv("DATABASE_URL")
    dbErr := database.InitDB(databaseURL)
    if dbErr != nil {
        panic(dbErr)
    }
    router := gin.Default()
    router.POST("/api/users", handlers.CreateUserHandler)
    router.POST("/api/users/generateotp", handlers.GenerateOTPHandler)
    router.POST("/api/users/verifyotp", handlers.VerifyOTPHandler)
    router.Run(":8080")
}
