package main

import (
	"encoding/json"
	"net/http"
)

func handleClientProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetClientProfile(w, r)
	case http.MethodPatch:
		UpdateClientProfile(w, r)
	case http.MethodDelete:
		DeleteClientProfile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetClientProfile(w http.ResponseWriter, r *http.Request) {
	clientid := r.URL.Query().Get("clientId")
	clientProfile, ok := database[clientid]
	if !ok || clientProfile == (ClientProfile{}) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := ClientProfile{
		Email: clientProfile.Email,
		Id:    clientProfile.Id,
		Name:  clientProfile.Name,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func UpdateClientProfile(w http.ResponseWriter, r *http.Request) {
	clientId := r.URL.Query().Get("clientId")
	if clientId == "" {
		http.Error(w, "clientId is required", http.StatusBadRequest)
		return
	}

	clientProfile, ok := database[clientId]
	if !ok {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}

	defer r.Body.Close()

	var payloadData ClientProfile
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Update only provided fields
	if payloadData.Name != "" {
		clientProfile.Name = payloadData.Name
	}
	if payloadData.Email != "" {
		clientProfile.Email = payloadData.Email
	}

	database[clientId] = clientProfile

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(clientProfile); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func DeleteClientProfile(w http.ResponseWriter, r *http.Request) {
	clientId := r.URL.Query().Get("clientId")
	if clientId == "" {
		http.Error(w, "clientId is required", http.StatusBadRequest)
		return
	}

	_, ok := database[clientId]
	if !ok {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}

	delete(database, clientId)
	w.WriteHeader(http.StatusNoContent)
}
