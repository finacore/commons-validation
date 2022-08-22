package commonsvalidation

import (
	errs "errors"

	commonserrors "github.com/finacore/commons-errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var Validate = validator.New()

//createTranslator function that obtain the translation
func createTranslator() ut.Translator {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	_ = en_translations.RegisterDefaultTranslations(Validate, trans)

	return trans
}

//createErrorArray function to create an CustomError array based on an
//error received by parameter
func createErrorArray(err error) []commonserrors.ValidationError {
	var errorArray []commonserrors.ValidationError

	var validationErrors validator.ValidationErrors
	errs.As(err, &validationErrors)

	for _, errElement := range validationErrors {
		translator := createTranslator()

		customError := commonserrors.CreateValidationError(
			errElement.StructNamespace(),
			errElement.Translate(translator),
		)

		// customError := newValidateError(errElement)
		errorArray = append(errorArray, *customError)
	}

	return errorArray
}

//Run function to perform the struct validation. This function returns an error
//array nil.
func ValidateModel(model interface{}) []commonserrors.ValidationError {
	if model == nil {
		return nil
	}

	err := Validate.Struct(model)

	if err != nil {
		array := createErrorArray(err)

		if len(array) > 0 {
			return array
		}
	}

	return nil
}
