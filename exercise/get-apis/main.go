package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"exercise/get-apis/handler"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	// dbHost = "postgres_db"
	// dbPort = "5432"
	// dbUser = "postgres"
	// dbPass = "postgres"
	// dbName = "postgres"

	dbHost = "localhost"
	dbPort = "5432"
	dbUser = "postgres"
	dbPass = "qwerty"
	dbName = "postgres"

	apiLogsTag = "get-apis: "
	Addr       = ":8080"
)

func main() {
	logger := log.New(os.Stdout, apiLogsTag, log.LstdFlags)
	// initializing database
	initDb(logger)
	defer db.Close()

	// configuring handlers
	h := handler.NewHandler(logger, db)
	sm := mux.NewRouter()
	getRouter := sm.PathPrefix("/api/v1/").Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/exercises-by/{id:[0-9]+}", h.GetExercisesByMuscleGroup)
	getRouter.HandleFunc("/muscle-groups-by/{id:[0-9]+}", h.GetMuscleGroupsByExercise)
	s := &http.Server{
		Addr:    Addr,
		Handler: sm,
	}
	err := s.ListenAndServe()
	if err != nil {
		logger.Fatalf("Error in ListenAndServe: Error: %v", err)
	}
}

// Initialize databse
func initDb(l *log.Logger) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)
	l.Print("Connecting to DB")
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		l.Panicf("Error in opening database connenction: %v", err)
	}
	// pinging db
	err = db.Ping()
	if err != nil {
		l.Panicf("Error in pinging database connenction: %v", err)
	}
	l.Print("Successfully DB connected!")
}
