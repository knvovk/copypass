package dto

type Status string

const (
	StatusOK    Status = "ok"
	StatusError Status = "error"
)

type Account struct {
	Id          string `json:"id,omitempty"`
	UserId      string `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

type User struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type Response struct {
	Status  Status
	Message string
	Data    any
}

func ErrorResponse(err error) Response {
	return Response{
		Status:  StatusError,
		Message: err.Error(),
		Data:    nil,
	}
}

func OkResponse(data any) Response {
	return Response{
		Status:  StatusOK,
		Message: "",
		Data:    data,
	}
}
