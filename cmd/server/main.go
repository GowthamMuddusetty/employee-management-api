package main

import (
	"log"

	"github.com/GowthamMuddusetty/employee-management-api/internal/config"
	"github.com/GowthamMuddusetty/employee-management-api/internal/db"
	"github.com/GowthamMuddusetty/employee-management-api/internal/handlers"
	"github.com/GowthamMuddusetty/employee-management-api/internal/middleware"
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

	employeeRepo := repositories.NewEmployeeRepository(postgresDB.Pool)
	employeeService := services.NewEmployeeService(employeeRepo)
	metricsRepo := repositories.NewSalaryMetricsRepository(postgresDB.Pool)

	userRepo := repositories.NewUserRepository(postgresDB.Pool)
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	authHandler := handlers.NewAuthHandler(authService)
	employeeHandler := handlers.NewEmployeeHandler(employeeService)
	salaryHandler := handlers.NewSalaryHandler(employeeService)
	metricsService := services.NewSalaryMetricsService(metricsRepo)
	metricsHandler := handlers.NewSalaryMetricsHandler(metricsService)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		employees := v1.Group("/employees")
		employees.Use(middleware.GinAuthMiddleware(cfg.JWTSecret))
		{
			employees.POST("", employeeHandler.Create)
			employees.GET("", employeeHandler.List)
			employees.GET("/:id", employeeHandler.GetByID)
			employees.PUT("/:id", employeeHandler.Update)
			employees.DELETE("/:id", employeeHandler.Delete)
			employees.GET("/:id/salary", salaryHandler.GetSalary)
		}

		metrics := v1.Group("/metrics")
		metrics.Use(middleware.GinAuthMiddleware(cfg.JWTSecret))
		{
			metrics.GET("/salary/country/:country", metricsHandler.ByCountry)
			metrics.GET("/salary/job-title/:jobTitle", metricsHandler.AvgByJobTitle)
		}
	}

	log.Println("server started on :8080")
	r.Run(":8080")
}
