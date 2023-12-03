package database

import (
	"aqarytest/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DbPool *pgxpool.Pool

// InitDB initializes the global database connection pool.
func InitDB(databaseURL string) error {
    var err error

    config, err := pgxpool.ParseConfig(databaseURL)
    if err != nil {
        return fmt.Errorf("error parsing database URL: %w", err)
    }

    DbPool, err = pgxpool.NewWithConfig(context.Background(), config)
    if err != nil {
        return fmt.Errorf("error creating database pool: %w", err)
    }

    fmt.Println("Connected to the database")
    return nil
}

// CreateUser inserts a new user into the database using a transaction.
func CreateUser(user models.User) error {
    ctx := context.Background()
    tx, err := DbPool.Begin(ctx)
    if err != nil {
        return err
    }
    defer tx.Rollback(ctx)

    const query = `INSERT INTO users (name, phone_number, otp, otp_created_at) VALUES ($1, $2, $3, $4) RETURNING id`
    if err := tx.QueryRow(ctx, query, user.Name, user.PhoneNumber, user.OTP, user.OTPCreatedAt).Scan(&user.ID); err != nil {
        return err
    }

    return tx.Commit(ctx)
}

// UpdateUserOTP updates the OTP for a user identified by their phone number using a transaction.
func UpdateUserOTP(phoneNumber, otp string) error {
    ctx := context.Background()
    tx, err := DbPool.Begin(ctx)
    if err != nil {
        return err
    }
    defer tx.Rollback(ctx)

    const query = `UPDATE users SET otp = $1, otp_created_at = NOW() WHERE phone_number = $2`
    if _, err := tx.Exec(ctx, query, otp, phoneNumber); err != nil {
        return err
    }

    return tx.Commit(ctx)
}

// GetUserByPhoneNumber retrieves a user by their phone number.
func GetUserByPhoneNumber(phoneNumber string) (*models.User, error) {
    const query = `SELECT id, name, phone_number, otp, otp_created_at FROM users WHERE phone_number = $1`
    user := &models.User{}

    err := DbPool.QueryRow(context.Background(), query, phoneNumber).Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.OTP, &user.OTPCreatedAt)
    if err != nil {
        return nil, err
    }

    return user, nil
}

// VerifyUserOTP verifies a user's OTP.
func VerifyUserOTP(phoneNumber, otp string) (bool, error) {
    var dbOTP string
    const query = `SELECT otp FROM users WHERE phone_number = $1`
    err := DbPool.QueryRow(context.Background(), query, phoneNumber).Scan(&dbOTP)
    if err != nil {
        return false, err
    }
    return dbOTP == otp, nil
}
