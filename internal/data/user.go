package data

type User struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}
