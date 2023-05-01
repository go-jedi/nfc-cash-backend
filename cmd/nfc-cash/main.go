package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	server "github.com/rob-bender/nfc-cash-backend"
	"github.com/rob-bender/nfc-cash-backend/pkg/handler"
	"github.com/rob-bender/nfc-cash-backend/pkg/repository"
	"github.com/rob-bender/nfc-cash-backend/pkg/service"
	"github.com/rob-bender/nfc-cash-backend/pkg/ws"
	"github.com/sirupsen/logrus"
)

//	@title			NFC-cash backend
//	@version		1.0
//	@description	API Server for NFC-cash Application

//	@host		localhost:8080
//	@BasePath	/

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	initConfig()
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "24972497Vlad",
		DBName:   "nfccash_db",
		SslMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	hub := ws.NewHub()
	handlers := handler.NewHandler(services, hub)
	go hub.Run()
	srv := new(server.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error accured while running http server in main.go: %s", err.Error())
	}
}

func initConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	for _, k := range []string{"PORT", "SMTP_EMAIL", "SMTP_PASSWORD"} {
		if _, ok := os.LookupEnv(k); !ok {
			log.Fatal("set environment variable -> ", k)
		}
	}
}
