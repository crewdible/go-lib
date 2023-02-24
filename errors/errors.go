package errors

import (
	"github.com/joomcode/errorx"
)

type (
	ErrorOption struct {
		HTTPCode int
	}
)

var (
	ErrHTTPCodeProperty = errorx.RegisterProperty("httpCode")
	ErrBase             = errorx.NewType(errorx.NewNamespace("crewdible"), "base")
)

func Errors(src error, opt *ErrorOption) error {
	err := ErrBase.New(src.Error())
	if opt == nil {
		return err
	}

	if opt.HTTPCode != 0 {
		err = err.WithProperty(ErrHTTPCodeProperty, opt.HTTPCode)
	}

	return err
}

func GetErrorHttpCode(err error) int {
	code, ok := errorx.ExtractProperty(err, ErrHTTPCodeProperty)
	if !ok {
		return 500
	}

	return code.(int)
}
