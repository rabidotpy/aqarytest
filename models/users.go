package models

import "time"

type User struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    PhoneNumber string    `json:"phone_number"`
    OTP         string    `json:"otp"`
    OTPCreatedAt time.Time `json:"otp_created_at"`
}

type OTPRequest struct {
    PhoneNumber string `json:"phone_number"`
}

type OTPVerificationRequest struct {
    PhoneNumber string `json:"phone_number"`
    OTP         string `json:"otp"`
}
