-- name: CreateUser :one
INSERT INTO users (name, phone_number, otp, otp_created_at)
VALUES ($1, $2, $3, $4)
RETURNING id, name, phone_number, otp, otp_created_at;

-- name: UpdateUserOTP :exec
UPDATE users
SET otp = $1, otp_created_at = NOW()
WHERE phone_number = $2;

-- name: GetUserByPhoneNumber :one
SELECT id, name, phone_number, otp, otp_created_at
FROM users
WHERE phone_number = $1;

-- name: VerifyUserOTP :one
SELECT otp FROM users WHERE phone_number = $1;
