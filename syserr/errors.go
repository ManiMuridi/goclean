package syserr

import "fmt"

type ValidationError map[string]string

func (v ValidationError) Error() string {
	var msg string

	for key, value := range v {
		msg += fmt.Sprintf("%s: %s\n", key, value)
	}

	return msg
}

type Error string

func (e Error) Error() string {
	return string(e)
}

func New(s string) error {
	return Error(s)
}
