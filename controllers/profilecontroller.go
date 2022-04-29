package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

const uri = "$MONGO_URI"

func GetProfiles(w http.ResponseWriter, r *http.Request) {
	// var profiles []entities.Profiles
	// database.Instance.Find(&profiles)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(profiles)
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	coll := client.Database("linktree").Collection("users")
	var result bson.M
	_ = coll.FindOne(context.TODO(), bson.D{{"username", "test"}}).Decode(&result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

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
