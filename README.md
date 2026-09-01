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

cmd/ Application entrypoints
configs/ Configuration files
global/ Shared application dependencies
internal/controller HTTP handlers
internal/initialize Application initialization
internal/middleware HTTP middleware
internal/router Route registration
pkg/database PostgreSQL helper
pkg/cache Redis helper
pkg/logger Logger setup
pkg/setting Configuration loader
migrations/ Database migrations

## Local Setup

Clone project:

```bash
git clone https://github.com/TanPhat-26/Go-ecommerce-backend-api.git
cd Go-ecommerce-backend-api
```
