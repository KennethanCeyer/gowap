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
