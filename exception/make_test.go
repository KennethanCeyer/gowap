package exception

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMake(t *testing.T) {
	var testError = Pair { "hello", 0x0001 }
	var makeError = Error(testError)

	assert.Equal(t, "exception(1): hello", makeError.Error())
}
