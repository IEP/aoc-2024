package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve01(t *testing.T) {
	input := "1.in"
	result, err := Solve01(input)
	assert.NoError(t, err)
	t.Logf("result:\n%s", result)
}

func TestSolve02(t *testing.T) {
	input := "1.in"
	result, err := Solve02(input)
	assert.NoError(t, err)
	t.Logf("result:\n%s", result)
}
