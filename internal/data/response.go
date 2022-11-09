package data

type Status string

const (
	StatusSuccess Status = "success"
	StatusFailure Status = "failure"
)

type Response struct {
	Status  Status
	Message string
	Data    any
}
