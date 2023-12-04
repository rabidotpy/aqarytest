package main

import (
	"aqarytest/database"
	"aqarytest/handlers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	questions "aqarytest/questions"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// we can use goose to manage migrations
// we can use docker to run the db
// we can use docker compose to run the app and the db

func main() {

    fmt.Println(questions.Questions2("akdcbkjbbbbbb"));
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

    authGroup := router.Group("/")
    authGroup.Use(AuthMiddleware())
    {
        // protected routes - 
        // remember i am using PhoneNumber in the header to authenticate user 
        // this is dummy authentication i am using for the sake of this test
        authGroup.POST("/api/users/generateotp", handlers.GenerateOTPHandler)
        authGroup.POST("/api/users/verifyotp", handlers.VerifyOTPHandler)
    }

    router.Run(":8080")
}

// authenticates based on the user pone number
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        phoneNumber := c.GetHeader("PhoneNumber")

        // chek if user exits in the db
        user, err := database.GetUserByPhoneNumber(context.Background(), phoneNumber)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized - user not found"})
            return
        }

        // attach user object to context or in node js terms, request object
        c.Set("user", user)

        c.Next() 
    }
}
