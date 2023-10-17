package main

import (
	"github.com/deuuus/bmsru-rsoi-lab2/src/reservation/internal/handlers"
	"github.com/deuuus/bmsru-rsoi-lab2/src/reservation/internal/repository"
	"github.com/deuuus/bmsru-rsoi-lab2/src/reservation/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	port := os.Getenv("PORT")

	db, err := repository.NewPostgresDB()
	if err != nil {
		logrus.Fatalf("error while db initializition: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handlers.NewHandler(service)
	router := handlers.InitRoutes()

	err = http.ListenAndServe(":"+port, router)

	if err != nil {
		logrus.Fatalf("Failed to start Reservation service on port %s", port)
	}
	logrus.Infof("Reservation service is listening on port %s", port)
}
