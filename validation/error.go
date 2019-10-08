package validation

import "fmt"

type Error map[string]string

func (v *Error) Error() string {
	var msg string

	for key, value := range *v {
		msg += fmt.Sprintf("%s: %s\n", key, value)
	}

	return msg
}
