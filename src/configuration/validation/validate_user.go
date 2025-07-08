package validation

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslation "github.com/go-playground/validator/v10/translations/en"
	"go-crud/src/configuration/rest_err"
	"log"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		enLocale := en.New()
		unt := ut.New(enLocale, enLocale)
		transl, _ = unt.GetTranslator("en")
		err := entranslation.RegisterDefaultTranslations(val, transl)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ValidateUserError(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationErr validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestErr("Invalid request body")
	} else if errors.As(validationErr, &jsonValidationErr) {
		var errorsCauses []rest_err.Causes

		for _, e := range validationErr.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}

		return rest_err.NewBadRequestValidationErr("Invalid request body", errorsCauses)
	}

	return rest_err.NewInternalServerErr(
		"Unexpected error occurred while validating user request")
}
