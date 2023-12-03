package handlers

import (
	"aqarytest/database"
	"aqarytest/models"
	"crypto/rand"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateOTPHandler(c *gin.Context) {
    var request models.OTPRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Find the user by phone number
    user, err := database.GetUserByPhoneNumber(request.PhoneNumber)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Generate an OTP
    otp, err := GenerateOTP(6) // Generate a 6-digit OTP
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate OTP"})
        return
    }

    // Update the user's OTP in the database
    if err := database.UpdateUserOTP(user.PhoneNumber, otp); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update OTP"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "OTP generated successfully"})
}

// GenerateOTP creates a cryptographically secure, random numeric OTP.
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

func VerifyOTPHandler(c *gin.Context) {
    var request models.OTPVerificationRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    valid, err := database.VerifyUserOTP(request.PhoneNumber, request.OTP)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify OTP"})
        return
    }

    if !valid {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired OTP"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "OTP verification successful"})
}
