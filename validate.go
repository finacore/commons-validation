// commonserrors package was created with intent to standardizing a way to execute the data
// structure validation processo that returns the validation erros every in the same way and
// in the same language.
//
// by convention, the english language was choosen as a default and unique language. Case your
// code want to use a diferente language the better for you is do not use this library and and
// choose to use the standard form of the go-playground/validator library
package commonsvalidation

import (
	errs "errors"

	commonserrors "github.com/finacore/commons-errors"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type costomValidator struct {
	validate   *validator.Validate
	translator ut.Translator

	validationErros []commonserrors.ValidationError
}

func New() *costomValidator {
	cv := &costomValidator{}

	cv.validate = validator.New()
	cv.translator = createTranslator(cv.validate)

	return cv
}

func (cv *costomValidator) ValidateModel(model interface{}) *costomValidator {
	if model == nil {
		return cv
	}

	err := cv.validate.Struct(model)

	if err != nil {
		cv.createErrorArray(err)
	}

	return cv
}

func (cv *costomValidator) HasError() bool { return len(cv.validationErros) > 0 }

func (cv *costomValidator) GetErrors() []commonserrors.ValidationError { return cv.validationErros }

func (cv *costomValidator) GetFirstError() *commonserrors.ValidationError {
	if cv.HasError() {
		return &cv.validationErros[0]
	}

	return nil
}

func (cv *costomValidator) createErrorArray(err error) {

	var validationErrors validator.ValidationErrors
	errs.As(err, &validationErrors)

	for _, errElement := range validationErrors {
		customError := commonserrors.CreateValidationError(
			errElement.StructNamespace(),
			errElement.Translate(cv.translator),
		)

		cv.validationErros = append(cv.validationErros, *customError)
	}
}
