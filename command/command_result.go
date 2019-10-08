package command

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type Result struct {
	Error error
	Data  interface{}
}

func (cr *Result) GetDataType() reflect.Type {
	return reflect.TypeOf(cr.Data)
}

func (cr *Result) Decode(data interface{}) {
	cr.Error = mapstructure.Decode(cr.Data, data)
}

func ErrorResult(err error) *Result {
	return &Result{
		Error: err,
		Data:  nil,
	}
}
