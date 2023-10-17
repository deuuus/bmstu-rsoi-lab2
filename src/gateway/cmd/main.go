package main

import (
	"github.com/deuuus/bmsru-rsoi-lab2/src/gateway/internal/handlers"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	port := os.Getenv("PORT")

	router := handlers.InitRoutes()

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		logrus.Fatalf("Failed to start API GateWay service on port %s", port)
	}
	logrus.Infof("API GateWay service is listening on port %s", port)
}
