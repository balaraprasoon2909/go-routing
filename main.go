package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/credentials", CredentialsHandler).Methods("POST")

	http.ListenAndServe(":8080", router)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	welcomeMessage := map[string]string{
		"message": "welcome to the golang back end dev project",
	}
	json.NewEncoder(w).Encode(welcomeMessage)
}

func CredentialsHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
		return
	}
	var credentialsApiRequest CredentialApiRequest
	err = json.Unmarshal(body, &credentialsApiRequest)
	fmt.Printf("api request body : %v", credentialsApiRequest)
	if err != nil {
		http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
		return
	}
	credentialsApiResponse := CredentialApiResponse{
		Username:      "test_usename",
		Password:      "test_password",
		Email:         "test_email",
		CustomerId:    "test_customer_id",
		AccountNumber: "account_number",
		Pin:           0,
	}
	err = json.NewEncoder(w).Encode(credentialsApiResponse)
	if err != nil {
		http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
		return
	}
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
