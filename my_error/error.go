package my_error

import "fmt"

type MyError interface {
	StatusCode() int32
	Message() string
	WithMessage(msg string) MyError
	Error() string
}

type myError struct {
	statusCode int32
	message    string
}

func (e *myError) StatusCode() int32 {
	return e.statusCode
}

func (e *myError) Message() string {
	return e.message
}

func (e *myError) WithMessage(msg string) MyError {
	res := e.clone()
	res.message = msg
	return res
}

func (e *myError) Error() string {
	return fmt.Sprintf("statusCode=%d, message=%s", e.statusCode, e.message)
}

func (e *myError) clone() *myError {
	res := &myError{
		statusCode: e.statusCode,
		message:    e.message,
	}
	return res
}
