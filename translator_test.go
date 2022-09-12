package commonsvalidation

import (
	"testing"

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
	vl := validator.New()
	tr := createTranslator(vl)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := vl.Struct(tt.args.model)
			vr := createErrorArray(err, tr)

			assert.NotNil(t, vr)
			assert.Equal(t, vr.Errors(), tt.want)
		})
	}
}
