// api/secrets.go
package api

import (
	"context"
	"encoding/json"
	"net/http"
	"database"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// Secret represents a secret entity
type Secret struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// PushSecret handles the creation of a new secret
func PushSecret(w http.ResponseWriter, r *http.Request) {
	var secret Secret
	err := json.NewDecoder(r.Body).Decode(&secret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.GetSecretsCollection().InsertOne(context.Background(), secret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetSecret retrieves a secret by its key
func GetSecret(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["key"]

	var secret Secret
	err := database.GetSecretsCollection().FindOne(context.Background(), bson.M{"key": key}).Decode(&secret)
	if err != nil {
		http.Error(w, "Secret not found", http.StatusNotFound)
		return
	}

	response := map[string]string{"key": secret.Key, "value": secret.Value}
	json.NewEncoder(w).Encode(response)
}

