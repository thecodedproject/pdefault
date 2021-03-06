// Code generated by pdefault/generated_tests/generator. DO NOT EDIT.

package generated_tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/thecodedproject/pdefault"
	"reflect"
	"testing"
)

func Test_Uint8_InitalisesFieldWhenNotSet(t *testing.T) {
	type TestStruct struct {
		Field *uint8 `pdefault:"10"`
	}

	input := TestStruct{}
	val := uint8(10)
	expected := TestStruct{
		Field: &val,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_Uint8_InitalisesMultipleFieldWhenNotSet(t *testing.T) {
	type TestStruct struct {
		Field1 *uint8 `pdefault:"10"`
		Field2 *uint8 `pdefault:"20"`
	}

	input := TestStruct{}
	val1 := uint8(10)
	val2 := uint8(20)
	expected := TestStruct{
		Field1: &val1,
		Field2: &val2,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_Uint8_InitalisesMultipleFieldWhenAllAreSet(t *testing.T) {
	type TestStruct struct {
		Field1 *uint8 `pdefault:"10"`
		Field2 *uint8 `pdefault:"20"`
	}

	val1 := uint8(30)
	val2 := uint8(30)
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

func Test_Uint8_InitalisesMultipleFieldWhenSomeAreSet(t *testing.T) {
	type TestStruct struct {
		Field1 *uint8 `pdefault:"10"`
		Field2 *uint8 `pdefault:"20"`
	}

	val1 := uint8(30)
	input := TestStruct{
		Field1: &val1,
	}
	val2 := uint8(20)
	expected := TestStruct{
		Field1: &val1,
		Field2: &val2,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_Uint8_NonPointerFieldsWithoutPdefaultAreIgnored(t *testing.T) {
	type TestStruct struct {
		Field1 *uint8 `pdefault:"10"`
		NonPointerField float64
	}

	input := TestStruct{
		NonPointerField: 1.0,
	}

	val := uint8(10)
	expected := TestStruct{
		Field1: &val,
		NonPointerField: 1.0,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_Uint8_WhenDefaultNotConvertibleToIntPanicsWhenFieldNil(t *testing.T) {

	// Do not perform this test on strings since a tag will always be convertable
	// to a string type.
	if reflect.TypeOf(uint8(10)).Kind() != reflect.String {
		type TestStruct struct {
			Field1 *uint8 `pdefault:"someString"`
		}

		input := TestStruct{
		}

		defer func() {
			if panicErr := recover(); panicErr != nil {
				assert.Regexp(t, "(.*someString)(.*uint8)(.*Field1)", panicErr, "Error should contain the field name")
			} else {
				assert.Fail(t, "Did not panic")
			}
		}()

		pdefault.Init(&input)
	}
}

func Test_Uint8_WhenDefaultNotConvertibleToIntPanicsWhenFieldNotNil(t *testing.T) {

	// Do not perform this test on strings since a tag will always be convertable
	// to a string type.
	if reflect.TypeOf(uint8(10)).Kind() != reflect.String {
		type TestStruct struct {
			Field1 *uint8 `pdefault:"hello"`
		}

		val := uint8(10)
		input := TestStruct{
			Field1: &val,
		}

		defer func() {
			if panicErr := recover(); panicErr != nil {
				assert.Regexp(t, "(.*hello)(.*uint8)(.*Field1)", panicErr, "Error should contain the field name")
			} else {
				assert.Fail(t, "Did not panic")
			}
		}()

		pdefault.Init(&input)
	}
}

func Test_Uint8_PdefaultTagOnNotPointerFieldPanics(t *testing.T) {

	type TestStruct struct {
		NonPointerField uint8 `pdefault:"10"`
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

func Test_Uint8_PanicsWithFieldNameIfDefaultValueOverflows(t *testing.T) {
	if "" != "" {
		assert.Fail(t, "some")

		type TestStruct struct {
			Field1 *uint8 `pdefault:""`
		}

		input := TestStruct{
		}

		defer func() {
			if panicErr := recover(); panicErr != nil {
				assert.Regexp(t, ".*Field1", panicErr, "Error should contain the field name")
			} else {
				assert.Fail(t, "Did not panic")
			}
		}()

		pdefault.Init(&input)

	}
}
