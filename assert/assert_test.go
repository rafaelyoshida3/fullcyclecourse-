package assert

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumum(t *testing.T) {
	got, err := Sum(2, 5)

	assert.Equal(t, 7, got, "they should not equal")
	assert.NotEqual(t, 6, got, "they should not be equal")
	assert.Nil(t, err)

	if err := errors.New("different then nil"); assert.NotNil(t, err) {
		assert.Equal(t, "different then nil", err.Error())
	}
}

// Assertion with table test
