package commonsvalidation

import (
	"reflect"
	"testing"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func Test_createTranslator(t *testing.T) {
	assert := assert.New(t)

	val := validator.New()
	trans := createTranslator(val)

	assert.NotNil(trans)
}

func Test_createErrorArray(t *testing.T) {
	type args struct {
		err        error
		translator ut.Translator
	}
	tests := []struct {
		name string
		args args
		want validatorResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createErrorArray(tt.args.err, tt.args.translator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createErrorArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
