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
