package lib

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPtr(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		got := Ptr("string")
		want := "string"
		assert.Equal(t, got, &want)
	})
	t.Run("bool", func(t *testing.T) {
		got := Ptr(true)
		want := true
		assert.Equal(t, got, &want)
	})
	t.Run("time", func(t *testing.T) {
		got := Ptr(time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC))
		want := time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, got, &want)
	})
	t.Run("int", func(t *testing.T) {
		got := Ptr(1234)
		want := 1234
		assert.Equal(t, got, &want)
	})
}
