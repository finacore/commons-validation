// commonserrors package was created with intent to standardizing a way to execute the data
// structure validation processo that returns the validation erros every in the same way and
// in the same language.
//
// by convention, the english language was choosen as a default and unique language. Case your
// code want to use a diferente language the better for you is do not use this library and and
// choose to use the standard form of the go-playground/validator library
package commonsvalidation

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// customValidator data structure that store the validator.Validate and ut.Translator objects, as
// well the list o of commonserrors.validationErros populated after execution of  Model method.
//
// This data structure is not public to avoid the programers create it directly, this way to create
// a instance of this struct is necessary call the New function.
type customValidator struct {
	validate   *validator.Validate
	translator ut.Translator
}

// New has the responsibility to create and instantiate the data structure that  able to create a pre
// convfigured validator that translate the errors to human english language.
func New() *customValidator {
	cv := &customValidator{}

	cv.validate = validator.New()
	cv.translator = createTranslator(cv.validate)

	return cv
}

// Model validator execute the validation of model passed by parameter. Once this method is called, the
// result is stored internaly in the data structure
func (cv *customValidator) Model(model interface{}) validatorResult {
	if model == nil {
		return validatorResult{}
	}

	err := cv.validate.Struct(model)

	if err != nil {
		return createErrorArray(err, cv.translator)
	}

	return validatorResult{}
}
