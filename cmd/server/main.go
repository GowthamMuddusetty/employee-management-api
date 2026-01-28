package main

import (
	"log"

	"github.com/GowthamMuddusetty/employee-management-api/internal/config"
	"github.com/GowthamMuddusetty/employee-management-api/internal/db"
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

	log.Println("employee-management-api started successfully")
}
