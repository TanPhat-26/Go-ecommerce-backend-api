# Go E-commerce Backend API

Backend REST API cho hệ thống thương mại điện tử, được xây dựng bằng Go và Gin.

## Tech Stack

- Go
- Gin
- PostgreSQL
- Redis
- GORM
- Viper
- Zap
- Docker Compose

## Requirements

- Go 1.26+
- Docker Desktop
- Git

## Project Structure

````text
Go-ecommerce-backend-api/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
│
├── configs/                        # Configuration templates
│
├── docs/                           # Project documentation
│   └── go-ecommerce-plan.md
│
├── global/
│   └── global.go                   # Shared application dependencies
│
├── internal/
│   ├── controller/                 # HTTP request handlers
│   │   └── health_controller.go
│   │
│   ├── initialize/                 # Application initialization
│   │   ├── config.go
│   │   ├── database.go
│   │   ├── logger.go
│   │   ├── redis.go
│   │   └── router.go
│   │
│   ├── middleware/                 # HTTP middleware
│   │   ├── access_log.go
│   │   ├── recovery.go
│   │   └── request_id.go
│   │
│   └── router/                     # Route registration
│       ├── router.go
│       └── router_test.go
│
├── migrations/                     # Database migration files
│
├── pkg/
│   ├── cache/                      # Redis client helper
│   │   └── redis.go
│   │
│   ├── database/                   # PostgreSQL connection helper
│   │   └── postgres.go
│   │
│   ├── logger/                     # Zap logger setup
│   │   ├── logger.go
│   │   └── logger_test.go
│   │
│   ├── setting/                    # Application configuration loader
│   │   ├── setting.go
│   │   ├── setting_test.go
│   │   └── testdata/
│   │       └── config.env
│   │
│   └── utils/                      # Shared utility functions
│       └── validator.go
│
├── response/                       # Standard API responses
├── scripts/                        # Development helper scripts
├── test/                           # Integration and end-to-end tests
├── third_party/                    # External service adapters
│
├── .env.example                    # Environment variables template
├── docker-compose.yml              # PostgreSQL and Redis services
├── go.mod                          # Go module definition
├── go.sum                          # Dependency checksums
└── README.md                       # Project documentation
'''

## Local Setup

Clone project:

```powershell
git clone https://github.com/TanPhat-26/Go-ecommerce-backend-api.git
cd Go-ecommerce-backend-api
````
