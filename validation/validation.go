package validation

import (
	"fmt"

	"github.com/ManiMuridi/goclean/translator"
	"gopkg.in/go-playground/validator.v9"
)

type Error struct {
	Field string
	Msg   string
	Tag   string
}

type Validator struct {
	Validator *validator.Validate
}

type ValidationError struct {
	Errors map[string]string
}

func (v *ValidationError) Error() string {
	var msg string

	for key, value := range v.Errors {
		msg += fmt.Sprintf("%s: %s\n", key, value)
	}

	return msg
}

func (v *Validator) Validate(i interface{}) error {
	errs := &ValidationError{Errors: make(map[string]string)}

	for _, e := range v.Validator.Struct(i).(validator.ValidationErrors) {
		msg := translator.T(e.Tag(), map[string]string{"Field": e.Field(), "Tag": e.Tag()}, nil)
		errs.Errors[e.Field()] = msg
	}

	return errs
}
