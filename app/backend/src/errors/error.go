package errors

import "fmt"

type MyError struct {
	Message string
	Code    int
}

func (m MyError) Error() string {
	return fmt.Sprint(m.Code) + " : " + m.Message
}
