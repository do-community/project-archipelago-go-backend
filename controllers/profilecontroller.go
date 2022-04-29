package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// func CreateProfile(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var profile entities.Profile
// 	json.NewDecoder(r.Body).Decode(&profile)
// 	database.Instance.Create(&profile)
// 	json.NewEncoder(w).Encode(entities.Profile)
// }

func GetProfileById(w http.ResponseWriter, r *http.Request) {
	profileId := mux.Vars(r)["id"]
	// if checkIfProfileExists(profileId) == false {
	// 	json.NewEncoder(w).Encode("Profile Not Found!")
	// 	return
	// }
	// var profile entities.Profile
	// database.Instance.First(&profile, profileId)
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(profile)
	fmt.Println(profileId)
}

func GetProfiles(w http.ResponseWriter, r *http.Request) {
	// var profiles []entities.Profiles
	// database.Instance.Find(&profiles)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(profiles)
	fmt.Println("hey, the GET route works")
}

// func UpdateProfile(w http.ResponseWriter, r *http.Request) {
// 	profileId := mux.Vars(r)["id"]
// 	if checkIfProfileExists(profileId) == false {
// 		json.NewEncoder(w).Encode("Profile Not Found!")
// 		return
// 	}
// 	var profile entities.Profile
// 	database.Instance.First(&profile, profileId)
// 	json.NewDecoder(r.Body).Decode(&profile)
// 	database.Instance.Save(&profile)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(profile)
// }

// func DeleteProfile(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	profileId := mux.Vars(r)["id"]
// 	if checkIfProfileExists(profileId) == false {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode("Profile Not Found!")
// 		return
// 	}
// 	var profile entities.Profile
// 	database.Instance.Delete(&profile, profileId)
// 	json.NewEncoder(w).Encode("Profile Deleted Successfully!")
// }

// func checkIfProfileExists(profileId string) bool {
// 	var profile entities.Profile
// 	database.Instance.First(&profile, profileId)
// 	if profile.ID == 0 {
// 		return false
// 	}
// 	return true
// }
