package errors

type MyError struct {
	Message string
	Code    int
}

func (m MyError) Error() string {
	return string(m.Code) + " : " + m.Message
}
