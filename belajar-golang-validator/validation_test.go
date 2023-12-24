package belajar_golang_validator

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidator(t *testing.T) {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		t.Error("validate is nil")
	}
}

func TestValidatorVariable(t *testing.T) {
	validate := validator.New()

	user := "e"

	err := validate.Var(user, "required")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidatorTwoVariable(t *testing.T) {
	validate := validator.New()

	password := "rahasia"
	confirmpassword := "rahasia"

	err := validate.VarWithValue(password, confirmpassword, "eqfield")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidatorMultiTag(t *testing.T) {
	validate := validator.New()

	age := "adaw"

	err := validate.Var(age, "required,numeric")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestStruct(t *testing.T) {
	type Login struct {
		Name     string `validate:"required,max=10"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()

	loginReq := Login{
		Name:     "1234567891",
		Password: "12345",
	}

	err := validate.Struct(loginReq)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func TestValidationError(t *testing.T) {
	type Login struct {
		Name     string `validate:"required,max=10"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()

	loginReq := Login{
		Name:     "1234567891",
		Password: "12345",
	}

	err := validate.Struct(loginReq)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func TestValidationErrors(t *testing.T) {
	type Login struct {
		Name     string `validate:"required,max=10"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()

	loginReq := Login{
		Name:     "12456227891",
		Password: "1345",
	}

	err := validate.Struct(loginReq)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationErr := range validationErrors {
			fmt.Println("error", validationErr.Field(), "on tag", validationErr.Tag(), "with err", validationErr.Error())
		}
	}

}
