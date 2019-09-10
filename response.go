package goclean

import (
	"strings"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/go-playground/validator.v9"
)

type Response interface {
	//service.BodyDecoder
}

type HttpResponse struct {
	Errors map[string]string
	Data   interface{}
}

func (h *HttpResponse) DecodeBody(obj interface{}) error {
	return mapstructure.Decode(h.Data, obj)
}

type response struct {
	Errors validator.ValidationErrorsTranslations
	Body   interface{}
}

func ValidationError(errString string) []string {
	return strings.Split(errString, "\n")
}

func (r *response) DecodeBody(obj interface{}) error {
	return mapstructure.Decode(r.Body, obj)
}

func NewResponse(body interface{}) Response {
	return &response{Body: body}
}

func ErrorResponse(err error) Response {
	//translation, _ := uni.GetTranslator("en")
	//
	//translator.GetTranslator()
	//validate.RegisterTranslation("required", translation, func(ut ut.Translator) error {
	//	return ut.Add("required", "{0} is required", true) // see universal-translator for details
	//}, func(ut ut.Translator, fe validator.FieldError) string {
	//	t, _ := ut.T("required", fe.Field())
	//
	//	return t
	//})

	//errs := err.(validator.ValidationErrors)
	return nil
	//res := &response{Body: nil, Errors: errs.Translate(translation)}
	//return &response{Body: nil, Errors: errs.Translate(translation)}
}
