package data

type Account struct {
	Id          string `json:"id,omitempty"`
	UserId      string `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}
