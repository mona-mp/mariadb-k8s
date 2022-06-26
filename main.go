package main

import (
	"log"
	"net/http"
	"os"

	"user-management/controllers"
	"user-management/database"
	"user-management/entity"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 18080")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":18080", router))

}

func initaliseHandlers(router *mux.Router) {

	router.HandleFunc("/users/{id}", controllers.UpdateUserByID).Methods("PUT")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.DeletUserByID).Methods("DELETE")
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")

}

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config :=
		database.Config{
			ServerName: os.Getenv("DB_SERVER"),
			User:       os.Getenv("DB_USER"),
			Password:   os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_DBNAME"),
		}
	connetionString := database.GetConnectionString(config)
	err = database.Connect(connetionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.User{})
}
