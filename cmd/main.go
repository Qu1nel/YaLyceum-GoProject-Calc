package main

import (
	"project/configs"
	"project/internal/calculator"
	"project/internal/errors"
	"project/internal/utils"
)

func main() {
	expression := errors.HandleError(func() (interface{}, error) {
		return utils.EnterExpression(configs.EnterExpressionPrompt)
	}).(string)

	result := errors.HandleError(func() (interface{}, error) {
		return calculator.Calc(expression)
	}).(float64)

	errors.HandleError(func() (interface{}, error) {
		return utils.PrintResultExpression(configs.LogMessage, result)
	})

}
