package validation

import (
	"github.com/ManiMuridi/goclean/syserr"
	"github.com/ManiMuridi/goclean/util"
	"gopkg.in/go-playground/validator.v9"
)

var Validator *V

func init() {
	Validator = &V{Validator: validator.New()}
}

type V struct {
	Validator *validator.Validate
}

func (v *V) Validate(i interface{}) error {
	vErrors := make(syserr.ValidationError)
	e := v.Validator.Struct(i)

	if e != nil {
		for _, e := range v.Validator.Struct(i).(validator.ValidationErrors) {
			msg := util.T(e.Tag(), map[string]string{"Field": e.Field(), "Tag": e.Tag()}, nil)
			vErrors[e.Field()] = msg
		}

		return vErrors
	}

	return nil
}
