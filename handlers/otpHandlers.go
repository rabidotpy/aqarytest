package handlers

import (
	"aqarytest/database"
	"aqarytest/models"
	"context"
	"crypto/rand"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)


func GenerateOTPHandler(c *gin.Context) {
    var request models.OTPRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx := context.Background()

    user, err := database.GetUserByPhoneNumber(ctx, request.PhoneNumber)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    otp, err := GenerateOTP(6) // 6 digits long otp
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate OTP"})
        return
    }

    updateParams := database.UpdateUserOTPParams{
        Otp:         pgtype.Text{String: otp},
        PhoneNumber: user.PhoneNumber,
    }
    if err := database.UpdateUserOTP(ctx, updateParams); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update OTP"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "OTP generated successfully", "otp": otp})
}


func VerifyOTPHandler(c *gin.Context) {
    var request models.OTPVerificationRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx := context.Background()

    dbOtp, err := database.VerifyUserOTP(ctx, request.PhoneNumber)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify OTP"})
        return
    }

    valid := dbOtp.String == request.OTP

    if !valid {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired OTP"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "OTP verification successful"})
}

func GenerateOTP(length int) (string, error) {
    const digits = "0123456789"
    otp := make([]byte, length)

    for i := range otp {
        num, err := rand.Int(rand.Reader, big.NewInt(10))
        if err != nil {
            return "", err
        }
        otp[i] = digits[num.Int64()]
    }

    return string(otp), nil
}
