// commonserrors package was created with intent to standardizing a way to execute the data
// structure validation processo that returns the validation erros every in the same way and
// in the same language.
//
// by convention, the english language was choosen as a default and unique language. Case your
// code want to use a diferente language the better for you is do not use this library and and
// choose to use the standard form of the go-playground/validator library
package commonsvalidation

import (
	"reflect"
	"testing"
)

func Test_validatorResult_Count(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vr := New().Model(tt.args.model)
			if got := vr.Count(); got != len(tt.want) {
				t.Errorf("validatorResult.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validatorResult_Errors(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vr := New().Model(tt.args.model)
			if got := vr.Errors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validatorResult.Errors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validatorResult_HasError(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vr := New().Model(tt.args.model)
			if got := vr.HasError(); got != (len(tt.want) > 0) {
				t.Errorf("validatorResult.HasError() = %v, want %v", got, tt.want)
			}
		})
	}
}
