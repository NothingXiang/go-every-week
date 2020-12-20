package vali

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// use single instance
var validate = validator.New()

func Check(in interface{}) error {

	err := validate.Struct(in)

	if err, ok := err.(*validator.InvalidValidationError); ok {
		fmt.Println(err)
		return err
	}

	if err != nil {
		//return fmt.Errorf("err:%w", err)
		return err
	}

	return nil
}
