package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Helper function to reset database to initial state
func resetDatabase() {
	database = map[string]ClientProfile{
		"user1": {
			Email: "email1@gmail.com",
			Id:    "user1",
			Name:  "User One",
			Token: "123",
		},
		"user2": {
			Email: "email2@gmail.com",
			Id:    "user2",
			Name:  "User Two",
			Token: "456",
		},
	}
}

func TestGetClientProfile_Success(t *testing.T) {
	resetDatabase()

	req := httptest.NewRequest(http.MethodGet, "/user/profile?clientId=user1", nil)
	w := httptest.NewRecorder()

	GetClientProfile(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response ClientProfile
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Email != "email1@gmail.com" {
		t.Errorf("Expected email 'email1@gmail.com', got '%s'", response.Email)
	}
	if response.Name != "User One" {
		t.Errorf("Expected name 'User One', got '%s'", response.Name)
	}
	if response.Id != "user1" {
		t.Errorf("Expected id 'user1', got '%s'", response.Id)
	}
}

func TestGetClientProfile_NotFound(t *testing.T) {
	resetDatabase()

	req := httptest.NewRequest(http.MethodGet, "/user/profile?clientId=nonexistent", nil)
	w := httptest.NewRecorder()

	GetClientProfile(w, req)

	if w.Code != http.StatusForbidden {
		t.Errorf("Expected status 403, got %d", w.Code)
	}
}

func TestGetClientProfile_MissingClientId(t *testing.T) {
	resetDatabase()

	req := httptest.NewRequest(http.MethodGet, "/user/profile", nil)
	w := httptest.NewRecorder()

	GetClientProfile(w, req)

	if w.Code != http.StatusForbidden {
		t.Errorf("Expected status 403, got %d", w.Code)
	}
}

func TestUpdateClientProfile_Success(t *testing.T) {
	resetDatabase()

	payload := map[string]string{
		"name":  "Updated User",
		"email": "updated@example.com",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPatch, "/user/profile?clientId=user1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	UpdateClientProfile(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response ClientProfile
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Name != "Updated User" {
		t.Errorf("Expected name 'Updated User', got '%s'", response.Name)
	}
	if response.Email != "updated@example.com" {
		t.Errorf("Expected email 'updated@example.com', got '%s'", response.Email)
	}

	// Verify database was updated
	if database["user1"].Name != "Updated User" {
		t.Error("Database was not updated")
	}
}

func TestUpdateClientProfile_PartialUpdate(t *testing.T) {
	resetDatabase()

	// Only update name, leave email unchanged
	payload := map[string]string{
		"name": "Only Name Changed",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPatch, "/user/profile?clientId=user2", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	UpdateClientProfile(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response ClientProfile
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Name != "Only Name Changed" {
		t.Errorf("Expected name 'Only Name Changed', got '%s'", response.Name)
	}
	// Email should remain unchanged
	if response.Email != "email2@gmail.com" {
		t.Errorf("Expected email 'email2@gmail.com', got '%s'", response.Email)
	}
}

func TestUpdateClientProfile_NotFound(t *testing.T) {
	resetDatabase()

	payload := map[string]string{
		"name": "Test",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPatch, "/user/profile?clientId=nonexistent", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	UpdateClientProfile(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestUpdateClientProfile_MissingClientId(t *testing.T) {
	resetDatabase()

	payload := map[string]string{
		"name": "Test",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPatch, "/user/profile", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	UpdateClientProfile(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestUpdateClientProfile_InvalidJSON(t *testing.T) {
	resetDatabase()

	req := httptest.NewRequest(http.MethodPatch, "/user/profile?clientId=user1", bytes.NewReader([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	UpdateClientProfile(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestDeleteClientProfile_Success(t *testing.T) {
	resetDatabase()

	req := httptest.NewRequest(http.MethodDelete, "/user/profile?clientId=user1", nil)
	w := httptest.NewRecorder()

	DeleteClientProfile(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", w.Code)
	}

	// Verify the client was deleted from database
	if _, exists := database["user1"]; exists {
		t.Error("Client was not deleted from database")
	}
}

func TestDeleteClientProfile_NotFound(t *testing.T) {
	resetDatabase()

	req := httptest.NewRequest(http.MethodDelete, "/user/profile?clientId=nonexistent", nil)
	w := httptest.NewRecorder()

	DeleteClientProfile(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestDeleteClientProfile_MissingClientId(t *testing.T) {
	resetDatabase()

	req := httptest.NewRequest(http.MethodDelete, "/user/profile", nil)
	w := httptest.NewRecorder()

	DeleteClientProfile(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestHandleClientProfile_MethodRouting(t *testing.T) {
	resetDatabase()

	tests := []struct {
		method       string
		expectedCode int
		description  string
	}{
		{http.MethodGet, http.StatusOK, "GET should succeed"},
		{http.MethodPatch, http.StatusBadRequest, "PATCH should fail without body"},
		{http.MethodDelete, http.StatusNoContent, "DELETE should succeed"},
		{http.MethodPost, http.StatusMethodNotAllowed, "POST should not be allowed"},
		{http.MethodPut, http.StatusMethodNotAllowed, "PUT should not be allowed"},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			resetDatabase()

			var req *http.Request
			if tt.method == http.MethodPatch {
				req = httptest.NewRequest(tt.method, "/user/profile", nil)
			} else {
				req = httptest.NewRequest(tt.method, "/user/profile?clientId=user1", nil)
			}

			w := httptest.NewRecorder()
			handleClientProfile(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("%s: Expected status %d, got %d", tt.description, tt.expectedCode, w.Code)
			}
		})
	}
}

func TestUpdateClientProfile_TokenNotModified(t *testing.T) {
	resetDatabase()

	originalToken := database["user1"].Token

	payload := map[string]string{
		"name":  "Updated Name",
		"token": "should-not-change",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPatch, "/user/profile?clientId=user1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	UpdateClientProfile(w, req)

	// Token should remain unchanged
	if database["user1"].Token != originalToken {
		t.Errorf("Token should not be modified. Expected '%s', got '%s'", originalToken, database["user1"].Token)
	}
}

func TestGetClientProfile_ContentType(t *testing.T) {
	resetDatabase()

	req := httptest.NewRequest(http.MethodGet, "/user/profile?clientId=user1", nil)
	w := httptest.NewRecorder()

	GetClientProfile(w, req)

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type 'application/json', got '%s'", contentType)
	}
}

func TestUpdateClientProfile_ContentType(t *testing.T) {
	resetDatabase()

	payload := map[string]string{
		"name": "Test",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPatch, "/user/profile?clientId=user1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	UpdateClientProfile(w, req)

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type 'application/json', got '%s'", contentType)
	}
}
