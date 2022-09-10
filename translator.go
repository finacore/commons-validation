// commonserrors package was created with intent to standardizing a way to execute the data
// structure validation processo that returns the validation erros every in the same way and
// in the same language.
//
// by convention, the english language was choosen as a default and unique language. Case your
// code want to use a diferente language the better for you is do not use this library and and
// choose to use the standard form of the go-playground/validator library
package commonsvalidation

import (
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
