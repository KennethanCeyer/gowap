/*
 * (C) 2017-2018 Kenneth Ceyer <kennethan@nhpcw.com>
 * this is distributed under
 * Apache 2.0 <https://www.apache.org/licenses/LICENSE-2.0>
 */

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
	"github.com/urfave/cli"
)

func setup(t *testing.T) func(t *testing.T) {
	fmt.Println("setup")
	app = initApp(cli.NewApp())
	return func(t *testing.T) {
		// TODO: add teardown
	}
}

func TestGowapApp(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)
	assert.NotZero(t, len(app.Commands))
}
