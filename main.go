package main

import (
	"archipelago-go-backend/controllers"
	"archipelago-go-backend/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database
	database.ConnectDB()
	// database.Connect(AppConfig.ConnectionString)
	// database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterProfileRoutes(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterProfileRoutes(router *mux.Router) {
	router.HandleFunc("/api/profiles", controllers.GetProfiles).Methods("GET")
	router.HandleFunc("/api/profiles/{id}", controllers.GetProfileById).Methods("GET")
	// router.HandleFunc("/api/profiles", controllers.CreateProfile).Methods("POST")
	// router.HandleFunc("/api/profiles/{id}", controllers.UpdateProfile).Methods("PUT")
	// router.HandleFunc("/api/profiles/{id}", controllers.DeleteProfile).Methods("DELETE")
}
