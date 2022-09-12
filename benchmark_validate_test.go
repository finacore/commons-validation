// commonserrors package was created with intent to standardizing a way to execute the data
// structure validation processo that returns the validation erros every in the same way and
// in the same language.
//
// by convention, the english language was choosen as a default and unique language. Case your
// code want to use a diferente language the better for you is do not use this library and and
// choose to use the standard form of the go-playground/validator library
package commonsvalidation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_New(b *testing.B) {
	vl := New()
	assert.NotNil(b, vl)
}

func Benchmark_Validate_Model(b *testing.B) {
	user := User{Name: "João", Surname: "da Silva", Email: "not-email"}

	cv := New()
	got := cv.Model(user)

	assert.Equal(b, got.Count(), 1)
}
