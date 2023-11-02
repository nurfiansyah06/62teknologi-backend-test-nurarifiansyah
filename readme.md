# Backend using Gin and Gorm 

Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API, but with performance up to 40 times faster than Martini. If you need smashing performance, get yourself some Gin.

## ⚡️ Quick start

1. Create a new project with Gin:

```bash
mkdir [folder-name] && cd [folder-name]
go mod init [folder-name]
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

2. Rename `.env.example` to `.env` and fill it with your environment values.

3. Go to API Docs page (Postman): [https://documenter.getpostman.com/view/11932880/2s9YXe63Uy](https://documenter.getpostman.com/view/11932880/2s9YXe63Uy)

## ⚙️ Configuration

```ini
# .env

# Stage status to start server:
#   - "dev", for start server without graceful shutdown
#   - "prod", for start server with graceful shutdown

# Database settings:
DB_HOST=
DB_DRIVER=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=

```

4. How to run the application:

```bash
go run main.go
```