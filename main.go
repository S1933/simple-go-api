package main

import (
"log"
"net/http"
)

func main() {
	http.HandleFunc("/user/profile", handleClientProfile)
	http.HandleFunc("/user/list", handleListClientProfiles)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
