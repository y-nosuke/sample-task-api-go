package presenter

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func BadRequestMessage(message string, err error) string {
	var invalidValidationError *validator.InvalidValidationError
	if errors.As(err, &invalidValidationError) {
		msg := invalidValidationError.Error()
		return msg
	}
	msg := message + "\n"
	msg += "errors: \n"
	for _, err := range err.(validator.ValidationErrors) {
		fmt.Println("=====================【validateチェック】===========================")
		fmt.Println("1: " + err.Namespace())
		fmt.Println("2: " + err.Field())
		fmt.Println("3: " + err.StructNamespace())
		fmt.Println("4: " + err.StructField())
		fmt.Println("5: " + err.Tag())
		fmt.Println("6: " + err.ActualTag())
		fmt.Println(err.Kind())
		fmt.Println(err.Type())
		fmt.Println(err.Value())
		fmt.Println("10" + err.Param())

		msg += "Namespace: " + err.Namespace() + "\n"
		msg += "Tag: " + err.Tag() + "\n"
		//message += "Value: " + err.Value().(string) + "\n"
	}
	fmt.Println("====================================================================")
	return msg
}
