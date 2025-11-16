package main

type ClientProfile struct {
	Email string
	Id    string
	Name  string
	Token string
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
