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

func BuildFailureResponse(err error) Response {
	return Response{
		Status:  StatusFailure,
		Message: err.Error(),
		Data:    nil,
	}
}

func BuildSuccessResponse(data any) Response {
	return Response{
		Status:  StatusSuccess,
		Message: "",
		Data:    data,
	}
}
