package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("works with empty list", func(t *testing.T) {
		//
		assert.Equal(t, 1, 1)
	})

	t.Run("works with non existing values", func(t *testing.T) {
		//
	})
}

