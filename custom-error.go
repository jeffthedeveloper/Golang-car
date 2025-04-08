package helpers

type CustomError struct {
	Message    string
	StatusCode int
	Method     string
	Err        error
}

func CustomErrBuilder(
	StatusCode int, Message, Method string, Err error,
) *CustomError {
	return &CustomError{Message, StatusCode, Method, Err}
}
