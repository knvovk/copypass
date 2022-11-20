package models

type User struct {
	Id           string
	Username     string
	Email        string
	PasswordHash string
}

type Account struct {
	Id          string
	User        User
	Name        string
	Description string
	Url         string
	Username    string
	Password    string
}
