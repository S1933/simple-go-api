package main

type ClientProfile struct {
	Email string
	Id    string
	Name  string
	Token string
}

type ClientListItem struct {
	Email string `json:"Email"`
	Id    string `json:"Id"`
	Name  string `json:"Name"`
}

type ClientListResponse struct {
	Clients    []ClientListItem `json:"clients"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	Total      int              `json:"total"`
	TotalPages int              `json:"total_pages"`
}

var database = map[string]ClientProfile{
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
