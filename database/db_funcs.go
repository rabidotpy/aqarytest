package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DbPool *pgxpool.Pool

//TODO: make use of goose for migrations

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

    fmt.Println("Yeyyy! Connected to the database")
    return nil
}

// TODO: make use of transaction as asked in assignment
// Rabi: using transaction here as a good practice only, we don't usually make use of transactions when creating user... 
func CreateUser(ctx context.Context, params CreateUserParams) (User, error) {
    tx, err := DbPool.Begin(ctx)
    if err != nil {
        return User{}, fmt.Errorf("error starting transaction: %w", err)
    }

    defer tx.Rollback(ctx)

    user, err := New(tx).CreateUser(ctx, params) // Using the transaction
    if err != nil {
        return User{}, fmt.Errorf("error creating user: %w", err)
    }

    if err := tx.Commit(ctx); err != nil {
        return User{}, fmt.Errorf("error committing transaction: %w", err)
    }

    return user, nil
}


func GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (User, error) {

    dduser, err := New(DbPool).GetUserByPhoneNumber(ctx, phoneNumber)
    fmt.Println("dduser: ", dduser);
    fmt.Println("err: ", err);
    if err != nil {
        return User{}, err
    }
    return dduser, nil
}

//TODO: make use of transaction
//i am making use of transaction here as a good practice only, not needed though
func UpdateUserOTP(ctx context.Context, params UpdateUserOTPParams) error {
    tx, err := DbPool.Begin(ctx)
    if err != nil {
        return err
    }

    //this anonymous function will be called when the function returns
    // it will check if there is an error, if there is an error it will rollback the transaction
    defer func() {
        if err != nil {
            tx.Rollback(ctx) 
        }
    }()

    if err = New(tx).UpdateUserOTP(ctx, params); err != nil {
        return err 
    }

    if err = tx.Commit(ctx); err != nil {
        return err
    }

    return nil
}



func VerifyUserOTP(ctx context.Context, phoneNumber string) (pgtype.Text, error) {
    return New(DbPool).VerifyUserOTP(ctx, phoneNumber)
}
