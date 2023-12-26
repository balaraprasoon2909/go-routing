package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main.go/mongo_client"
)

var client *mongo.Client

func main() {
	client = mongo_client.Client
	defer mongo_client.CloseMongoClient()
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/credentials", GetCredentials).Methods("POST")

	http.ListenAndServe(":8080", router)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

func GetCredentials(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
		return
	}
	var credentialsApiRequest CredentialApiRequest
	err = json.Unmarshal(body, &credentialsApiRequest)
	fmt.Printf("api request body : %v\n", credentialsApiRequest)
	if err != nil {
		http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
		return
	}
	response, err := getInfoForId(credentialsApiRequest.Name, credentialsApiRequest.Password)
	if err != nil {
		fmt.Printf("Received an error while getting information from DB : %+v", err)
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
		return
	}
}

func getInfoForId(id string, password string) (CredentialApiResponse, error) {
	if client == nil {
		return CredentialApiResponse{}, errors.New("empty mongo client")
	}
	var response CredentialApiResponse
	filter := bson.D{{Key: "username", Value: id}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := client.Database("test").Collection("credentials").Find(ctx, filter)
	if err != nil {
		return response, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		if err := cursor.Decode(&response); err != nil {
			return response, err
		}
	}
	return response, nil
}

type CredentialApiResponse struct {
	Username      string `json:"username,omitempty"`
	Password      string `json:"password,omitempty"`
	Email         string `json:"email,omitempty"`
	CustomerId    string `json:"customer_id,omitempty"`
	AccountNumber string `json:"accountNumber,omitempty"`
	Pin           int64  `json:"pin,omitempty"`
}

type CredentialApiRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
