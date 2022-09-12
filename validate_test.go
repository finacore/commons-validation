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

	commonserrors "github.com/finacore/commons-errors"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Name    string `validate:"required,min=3,max=256"`
	Surname string `validate:"required,min=3,max=256"`
	Email   string `validate:"omitempty,email"`
}

type args struct {
	model interface{}
}

var tests = []struct {
	name string
	args args
	want []commonserrors.ValidationError
}{
	{
		name: "nil-model",
		args: args{model: nil},
		want: []commonserrors.ValidationError{},
	},
	{
		name: "all-in",
		args: args{model: User{Name: "João", Surname: "da Silva", Email: "dasilva@gmail.com"}},
		want: []commonserrors.ValidationError{},
	},
	{
		name: "no-name",
		args: args{model: User{Surname: "da Silva", Email: "dasilva@gmail.com"}},
		want: []commonserrors.ValidationError{{Field: "User.Name", Message: "Name is a required field"}},
	},
	{
		name: "no-surname",
		args: args{model: User{Name: "João", Email: "dasilva@gmail.com"}},
		want: []commonserrors.ValidationError{{Field: "User.Surname", Message: "Surname is a required field"}},
	},
	{
		name: "no-email",
		args: args{model: User{Name: "João", Surname: "da Silva"}},
		want: []commonserrors.ValidationError{},
	},
	{
		name: "invalid-email",
		args: args{model: User{Name: "João", Surname: "da Silva", Email: "not-email"}},
		want: []commonserrors.ValidationError{{Field: "User.Email", Message: "Email must be a valid email address"}},
	},
	{
		name: "invalid-email-&-short-name",
		args: args{model: User{Name: "Jo", Surname: "da Silva", Email: "not-email"}},
		want: []commonserrors.ValidationError{
			{
				Field:   "User.Name",
				Message: "Name must be at least 3 characters in length",
			},
			{
				Field:   "User.Email",
				Message: "Email must be a valid email address",
			},
		},
	},
}

func TestNew(t *testing.T) {
	assert := assert.New(t)

	vl := New()
	assert.NotNil(vl)

	assert.NotNil(vl.validate)
	assert.NotNil(vl.translator)
}

// func TestModel(t *testing.T) {
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			vl := New()
// 			vl.Model(tt.args.model)
// 			assert.Equal(t, vl.validationErros, tt.want)
// 		})
// 	}
// }

// func TestHasError(t *testing.T) {
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			vl := New()
// 			hasError := vl.Model(tt.args.model).HasError()

// 			assert.Equal(t, len(tt.want) > 0, hasError)
// 		})
// 	}
// }

// func TestErrors(t *testing.T) {
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			vl := New()
// 			errs := vl.Model(tt.args.model).Errors()

// 			assert.Equal(t, tt.want, errs)
// 		})
// 	}
// }

// func TestFirstError(t *testing.T) {
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			vl := New()
// 			err := vl.Model(tt.args.model).FirstError()

// 			assert.Equal(t, tt.want, vl.Errors())
// 			if len(tt.want) > 0 {
// 				assert.Equal(t, tt.want[0], *err)
// 			} else {
// 				assert.Nil(t, err)
// 			}
// 		})
// 	}
// }

func Test_customValidator_Model(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cv := New()

			got := cv.Model(tt.args.model)
			assert.NotNil(t, got)

			assert.Equal(t, got.Errors(), tt.want)
		})
	}
}
