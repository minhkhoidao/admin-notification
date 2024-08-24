package server

import (
	"admin-backend/database"
	"admin-backend/routes"
	"log"
	"net/http"
	"time"
)

var timeDuration = 30 * time.Second

func NewServer() *http.Server {
	dbInstance := database.NewConnection()

	if err := dbInstance.Migrate(); err != nil {
		log.Println("Error migrating database: ", err)
	}
	server := &http.Server{
		Addr:         ":8080",
		Handler:      routes.Setup(timeDuration, dbInstance.GetDB()),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return server
}
