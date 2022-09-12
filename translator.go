// commonserrors package was created with intent to standardizing a way to execute the data
// structure validation processo that returns the validation erros every in the same way and
// in the same language.
//
// by convention, the english language was choosen as a default and unique language. Case your
// code want to use a diferente language the better for you is do not use this library and and
// choose to use the standard form of the go-playground/validator library
//
// The first step to use this package is to get it through the command below, then you will be
// able to import it in your code.
//
//	go get github.com/finacore/commons-validation
//
// To avoid the conflicts between packages during the imnport phase, is strongly recommended
// select other name for this package in the moment of import like shown below:
//
//	import (
//		commonsvalidation "github.com/finacore/commons-validation"
//	)
//
// Once imported you can use the package as you prefer. So read the function and methods
// documentation to know how to use this packagage.
package commonsvalidation

import (
	"errors"

	commonserrors "github.com/finacore/commons-errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// createTranslator is responsible to create a ut.Translator object and register itself as the
// translator for validator.Validate. By definition the language selected is english, so if you
// want to use other language is recomended do not use this package or fork it and change this
// function
//
// For now, has no plans or project to make it multi language. If you want to add suport to multi
// language you are wellcome to make a pull request. Thanks
func createTranslator(validate *validator.Validate) ut.Translator {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	return trans
}

// createErrorArray has the responsibility to create the arror array based on the response of validation
// and the given translator.
func createErrorArray(err error, translator ut.Translator) ValidatorResult {
	vr := ValidatorResult{}

	//transfor the erros to
	var validationErrors validator.ValidationErrors
	errors.As(err, &validationErrors)

	for _, errElement := range validationErrors {
		customError := commonserrors.CreateValidationError(
			errElement.StructNamespace(),
			errElement.Translate(translator),
		)

		vr = append(vr, *customError)
	}

	return vr
}
