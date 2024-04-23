package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecution(t *testing.T) {
	t.Run("returns 0", func(t *testing.T) {
		assert.NotPanics(t, main)
	})
}
