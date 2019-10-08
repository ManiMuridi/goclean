package validation

import (
	"github.com/ManiMuridi/goclean/util"
	"gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	Validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	vErrors := Error{}
	e := v.Validator.Struct(i)

	if e != nil {
		for _, e := range v.Validator.Struct(i).(validator.ValidationErrors) {
			msg := util.T(e.Tag(), map[string]string{"Field": e.Field(), "Tag": e.Tag()}, nil)
			vErrors[e.Field()] = msg
		}
	}

	return &vErrors
}
