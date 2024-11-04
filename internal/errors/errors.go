package errors

import (
	stdErr "errors"
	"log"
	"project/configs"
)

var (
	ErrDivisionByZero    = stdErr.New("деление на ноль")
	ErrEmptyExpression   = stdErr.New("пустое выражение")
	ErrMissingBracket    = stdErr.New("пропущена скобка")
	ErrInvalidExpression = stdErr.New("неккоректное выражение")
	ErrNotEnogthOperand  = stdErr.New("не достаточно операндов")
)

func HandleError(fn func() (interface{}, error)) interface{} {
	result, err := fn()
	if err != nil {
		log.Fatalf(configs.ErrorMsg+" %s\n", err.Error())
	}
	return result
}
