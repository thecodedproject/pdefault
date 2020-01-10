package pdefault_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/thecodedproject/pdefault"
	"testing"
)

func TestSome(t *testing.T) {
	type SomeType struct {
		SomeInteger        *int64   `pdefault:"10"`
		SomeFloat          *float64 `pdefault:"1.5"`
		SomeNonPointerType string
	}

	s := SomeType{
		SomeNonPointerType: "Hello World!",
	}

	pdefault.Init(&s)

	i := int64(10)
	f := float64(1.5)
	expected := SomeType{
		SomeInteger:        &i,
		SomeFloat:          &f,
		SomeNonPointerType: "Hello World!",
	}

	assert.Equal(t, expected, s)
}
