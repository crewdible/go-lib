package filter

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type ErrorValidateResponse struct {
	FailedField string `json:"failed_field,omitempty"`
	Tag         string `json:"tag,omitempty"`
	Value       string `json:"value,omitempty"`
}

var validate = validator.New()

func Validate(s interface{}) []*ErrorValidateResponse {
	var errors []*ErrorValidateResponse
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorValidateResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func Empty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.String:
		return v.String() == ""
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Interface, reflect.Chan:
		return v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	}
	return false
}

func SetField(obj interface{}, name string, value interface{}) error {
	// Fetch the field reflect.Value
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	// If obj field value is not settable an error is thrown
	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		invalidTypeError := errors.New("Provided value type didn't match obj field type")
		return invalidTypeError
	}

	structFieldValue.Set(val)
	return nil
}
