package goclean

import (
	"github.com/hashicorp/go-multierror"
	"github.com/labstack/echo/v4"
)

type Request interface {
	Decode(interface{}) error
	//Get(name string) interface{}
	//Error() error
}

type request struct {
	data    interface{}
	params  interface{}
	err     *multierror.Error
	context echo.Context
}

func NewRequest(data interface{}) Request {
	return &request{data: data}
}

//func (r *request) Error() error {
//	return r.err.ErrorOrNil()
//}

//
//func (r *request) Get(name string) interface{} {
//	if str := r.context.Param(name); str != "" {
//		r.data = str
//	} else if str := r.context.QueryParam(name); str != "" {
//		r.data = str
//	} else if str := r.context.FormValue(name); str != "" {
//		r.data = str
//	} else if file, err := r.context.FormFile(name); file != nil {
//		r.err = multierror.Append(r.err, err)
//		r.data = file
//	}
//
//	return r.data
//}

func (r *request) Decode(obj interface{}) error {
	return r.context.Bind(&obj)
}

func NewEchoRequest(c echo.Context) Request {
	//paramNames := c.ParamNames()
	//paramValues := c.ParamValues()
	//paramVars := make(url.Values)
	//
	//for in, name := range paramNames {
	//	paramVars[name] = append(paramVars[name], paramValues[in])
	//}

	return &request{
		context: c,
	}
}
