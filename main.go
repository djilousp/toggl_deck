package main

import (
	"net/http"
	"os"
	"time"

	"github.com/djilousp/toggl_deck/src/infrastructure/persistence"
	"github.com/gorilla/mux"
)

func main() {

	dbdriver := "postgres"
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	
	services, err := persistence.NewRepositories(dbdriver, user, password, port, host, dbname)
	if err != nil {
		panic(err)
	}
	defer services.Close()
	services.Automigrate()

	r := mux.NewRouter()

	s := http.Server{
		Addr:         ":9999",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	s.ListenAndServe()
}