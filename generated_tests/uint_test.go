// Code generated by pdefault/generated_tests/generator. DO NOT EDIT.

package generated_tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/thecodedproject/pdefault"
	"reflect"
	"testing"
)

func Test_Uint_InitalisesFieldWhenNotSet(t *testing.T) {
	type TestStruct struct {
		Field *uint `pdefault:"10"`
	}

	input := TestStruct{}
	val := uint(10)
	expected := TestStruct{
		Field: &val,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_Uint_InitalisesMultipleFieldWhenNotSet(t *testing.T) {
	type TestStruct struct {
		Field1 *uint `pdefault:"10"`
		Field2 *uint `pdefault:"20"`
	}

	input := TestStruct{}
	val1 := uint(10)
	val2 := uint(20)
	expected := TestStruct{
		Field1: &val1,
		Field2: &val2,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_Uint_InitalisesMultipleFieldWhenAllAreSet(t *testing.T) {
	type TestStruct struct {
		Field1 *uint `pdefault:"10"`
		Field2 *uint `pdefault:"20"`
	}

	val1 := uint(30)
	val2 := uint(30)
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

func Test_Uint_InitalisesMultipleFieldWhenSomeAreSet(t *testing.T) {
	type TestStruct struct {
		Field1 *uint `pdefault:"10"`
		Field2 *uint `pdefault:"20"`
	}

	val1 := uint(30)
	input := TestStruct{
		Field1: &val1,
	}
	val2 := uint(20)
	expected := TestStruct{
		Field1: &val1,
		Field2: &val2,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_Uint_NonPointerFieldsWithoutPdefaultAreIgnored(t *testing.T) {
	type TestStruct struct {
		Field1 *uint `pdefault:"10"`
		NonPointerField float64
	}

	input := TestStruct{
		NonPointerField: 1.0,
	}

	val := uint(10)
	expected := TestStruct{
		Field1: &val,
		NonPointerField: 1.0,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_Uint_WhenDefaultNotConvertibleToIntPanicsWhenFieldNil(t *testing.T) {

	// Do not perform this test on strings since a tag will always be convertable
	// to a string type.
	if reflect.TypeOf(uint(10)).Kind() != reflect.String {
		type TestStruct struct {
			Field1 *uint `pdefault:"someString"`
		}

		input := TestStruct{
		}

		defer func() {
			if panicErr := recover(); panicErr != nil {
				assert.Regexp(t, "(.*someString)(.*uint)(.*Field1)", panicErr, "Error should contain the field name")
			} else {
				assert.Fail(t, "Did not panic")
			}
		}()

		pdefault.Init(&input)
	}
}

func Test_Uint_WhenDefaultNotConvertibleToIntPanicsWhenFieldNotNil(t *testing.T) {

	// Do not perform this test on strings since a tag will always be convertable
	// to a string type.
	if reflect.TypeOf(uint(10)).Kind() != reflect.String {
		type TestStruct struct {
			Field1 *uint `pdefault:"hello"`
		}

		val := uint(10)
		input := TestStruct{
			Field1: &val,
		}

		defer func() {
			if panicErr := recover(); panicErr != nil {
				assert.Regexp(t, "(.*hello)(.*uint)(.*Field1)", panicErr, "Error should contain the field name")
			} else {
				assert.Fail(t, "Did not panic")
			}
		}()

		pdefault.Init(&input)
	}
}

func Test_Uint_PdefaultTagOnNotPointerFieldPanics(t *testing.T) {

	type TestStruct struct {
		NonPointerField uint `pdefault:"10"`
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

func Test_Uint_PanicsWithFieldNameIfDefaultValueOverflows(t *testing.T) {
	if "" != "" {
		assert.Fail(t, "some")

		type TestStruct struct {
			Field1 *uint `pdefault:""`
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
