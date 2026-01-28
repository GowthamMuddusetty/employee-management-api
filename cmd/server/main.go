package main

import (
	"log"

	"github.com/GowthamMuddusetty/employee-management-api/internal/config"
	"github.com/GowthamMuddusetty/employee-management-api/internal/db"
	"github.com/GowthamMuddusetty/employee-management-api/internal/handlers"
	"github.com/GowthamMuddusetty/employee-management-api/internal/repositories"
	"github.com/GowthamMuddusetty/employee-management-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	postgresDB, err := db.New(cfg)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer postgresDB.Pool.Close()

	userRepo := repositories.NewUserRepository(postgresDB.Pool)
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	authHandler := handlers.NewAuthHandler(authService)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}
	}

	log.Println("server started on :8080")
	r.Run(":8080")
}
