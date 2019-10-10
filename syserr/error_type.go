package syserr

import "reflect"

const (
	Validation = "validation.Error"
)

func Type(err error) string {
	if err != nil {
		return reflect.TypeOf(err).String()
	}
	return ""
}
