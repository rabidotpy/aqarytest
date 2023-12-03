package handlers

import (
	"aqarytest/database"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)



func CreateUserHandler(c *gin.Context) {
    var request database.CreateUserRequest
    ctx := context.Background()
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := database.GetUserByPhoneNumber(ctx, request.PhoneNumber)
    if err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
        return
    }

    if err != pgx.ErrNoRows {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check if user exists"})
        return
    }

    generatedOtp, err := GenerateOTP(6)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate OTP"})
        return
    }

    createUserParams := database.CreateUserParams{
        Name:         request.Name,
        PhoneNumber:  request.PhoneNumber,
        Otp:          pgtype.Text{String: generatedOtp},
        OtpCreatedAt: pgtype.Timestamptz{Time: time.Now()},
    }

    returnedUser, err := database.CreateUser(ctx, createUserParams)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": returnedUser})
}





// func CreateUserHandler(c *gin.Context) {
//     var newUser database.CreateUserParams
//     ctx := context.Background()
//     if err := c.ShouldBindJSON(&newUser); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }
   
//     generatedOtp, err := GenerateOTP(6)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate OTP"})
//         return
//     }
//     fmt.Println((generatedOtp), "generatedOtp")
//     newUser.Otp = pgtype.Text{String: generatedOtp}
// newUser.OtpCreatedAt = pgtype.Timestamptz{Time: time.Now()}

//     fmt.Println((newUser), "newUser")
//     returnedUser, err := database.CreateUser(ctx, newUser)
//     if err != nil {
//         fmt.Println((err), "err")
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
//         return
//     }
//     c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": returnedUser})
// }
