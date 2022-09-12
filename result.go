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

import commonserrors "github.com/finacore/commons-errors"

// validatorResult is da data structure that has a responsibility to store the response of the
// validation process. This data structure implements some methods to help get access to the data
type validatorResult []commonserrors.ValidationError

// Count returns how many validation errors was produced by the validation process execution
//
// Usage:
//
//	cv := commonsvalidation.New()
//	vr := cv.Model(yourModel)
//	numErrors := vr.Count()
func (vr validatorResult) Count() int { return len(vr) }

// HasError provivide a simple way to check if the validatorResponse produce or not errors
//
// Usage:
//
//	cv := commonsvalidation.New()
//	vr := cv.Model(yourModel)
//	hasError := vr.HasError()
func (vr validatorResult) HasError() bool { return vr.Count() > 0 }

// Errors return the commonserrors.ValidationError array containing all errors returned by the
// execution of validation process
//
// Usage:
//
//	cv := commonsvalidation.New()
//	vr := cv.Model(yourModel)
//	errs := vr.Errors()
func (vr validatorResult) Errors() []commonserrors.ValidationError { return vr }
