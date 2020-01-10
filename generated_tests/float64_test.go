package generated_tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/thecodedproject/pdefault"
	"reflect"
	"testing"
)

func Test_Float64_InitalisesFieldWhenNotSet(t *testing.T) {
	type TestStruct struct {
		Field *float64 `pdefault:"1.1"`
	}

	input := TestStruct{}
	val := float64(1.1)
	expected := TestStruct{
		Field: &val,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_Float64_InitalisesMultipleFieldWhenNotSet(t *testing.T) {
	type TestStruct struct {
		Field1 *float64 `pdefault:"1.1"`
		Field2 *float64 `pdefault:"2.2"`
	}

	input := TestStruct{}
	val1 := float64(1.1)
	val2 := float64(2.2)
	expected := TestStruct{
		Field1: &val1,
		Field2: &val2,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_Float64_InitalisesMultipleFieldWhenAllAreSet(t *testing.T) {
	type TestStruct struct {
		Field1 *float64 `pdefault:"1.1"`
		Field2 *float64 `pdefault:"2.2"`
	}

	val1 := float64(3.3)
	val2 := float64(3.3)
	input := TestStruct{
		Field1: &val1,
		Field2: &val2,
	}
	expected := TestStruct{
		Field1: &val1,
		Field2: &val2,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_Float64_InitalisesMultipleFieldWhenSomeAreSet(t *testing.T) {
	type TestStruct struct {
		Field1 *float64 `pdefault:"1.1"`
		Field2 *float64 `pdefault:"2.2"`
	}

	val1 := float64(3.3)
	input := TestStruct{
		Field1: &val1,
	}
	val2 := float64(2.2)
	expected := TestStruct{
		Field1: &val1,
		Field2: &val2,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_Float64_NonPointerFieldsWithoutPdefaultAreIgnored(t *testing.T) {
	type TestStruct struct {
		Field1 *float64 `pdefault:"1.1"`
		NonPointerField float64
	}

	input := TestStruct{
		NonPointerField: 1.0,
	}

	val := float64(1.1)
	expected := TestStruct{
		Field1: &val,
		NonPointerField: 1.0,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_Float64_WhenDefaultNotConvertibleToIntPanicsWhenFieldNil(t *testing.T) {

	// Do not perform this test on strings since a tag will always be convertable
	// to a string type.
	if reflect.TypeOf(float64(1.1)).Kind() != reflect.String {
		type TestStruct struct {
			Field1 *float64 `pdefault:"someString"`
		}

		input := TestStruct{
		}

		defer func() {
			if panicErr := recover(); panicErr != nil {
				assert.Regexp(t, "(.*someString)(.*float64)(.*Field1)", panicErr, "Error should contain the field name")
			} else {
				assert.Fail(t, "Did not panic")
			}
		}()

		pdefault.Init(&input)
	}
}

func Test_Float64_WhenDefaultNotConvertibleToIntPanicsWhenFieldNotNil(t *testing.T) {

	// Do not perform this test on strings since a tag will always be convertable
	// to a string type.
	if reflect.TypeOf(float64(1.1)).Kind() != reflect.String {
		type TestStruct struct {
			Field1 *float64 `pdefault:"hello"`
		}

		val := float64(1.1)
		input := TestStruct{
			Field1: &val,
		}

		defer func() {
			if panicErr := recover(); panicErr != nil {
				assert.Regexp(t, "(.*hello)(.*float64)(.*Field1)", panicErr, "Error should contain the field name")
			} else {
				assert.Fail(t, "Did not panic")
			}
		}()

		pdefault.Init(&input)
	}
}

func Test_Float64_PdefaultTagOnNotPointerFieldPanics(t *testing.T) {

	type TestStruct struct {
		NonPointerField float64 `pdefault:"1.1"`
	}

	input := TestStruct{
	}

	defer func() {
		if panicErr := recover(); panicErr != nil {
			assert.Regexp(t, "NonPointerField", panicErr, "Error should contain the field name")
		} else {
			assert.Fail(t, "Did not panic")
		}
	}()

	pdefault.Init(&input)
}
