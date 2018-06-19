/*
 * (C) 2017-2018 Kenneth Ceyer <kennethan@nhpcw.com>
 * this is distributed under
 * Apache 2.0 <https://www.apache.org/licenses/LICENSE-2.0>
 */

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
