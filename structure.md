aqarytest/
│
├── main.go
├── go.mod
├── go.sum
│
├── sqlc/ # Folder for SQL files and sqlc configuration
│ ├── schema.sql # SQL schema definitions
│ ├── queries.sql # SQL query definitions
│ └── sqlc.yaml # sqlc configuration file
│
├── models/ # Go models
│ └── user.go
│
├── database/ # sqlc generated code and database connection logic
│ └── db.go
│
└── handlers/ # Handlers for different endpoints
├── userHandlers.go
└── otpHandlers.go
