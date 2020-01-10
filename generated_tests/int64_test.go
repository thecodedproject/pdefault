package generated_tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/thecodedproject/pdefault"
	"reflect"
	"testing"
)

func Test_Int64_InitalisesFieldWhenNotSet(t *testing.T) {
	type TestStruct struct {
		Field *int64 `pdefault:"10"`
	}

	input := TestStruct{}
	val := int64(10)
	expected := TestStruct{
		Field: &val,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_Int64_InitalisesMultipleFieldWhenNotSet(t *testing.T) {
	type TestStruct struct {
		Field1 *int64 `pdefault:"10"`
		Field2 *int64 `pdefault:"20"`
	}

	input := TestStruct{}
	val1 := int64(10)
	val2 := int64(20)
	expected := TestStruct{
		Field1: &val1,
		Field2: &val2,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_Int64_InitalisesMultipleFieldWhenAllAreSet(t *testing.T) {
	type TestStruct struct {
		Field1 *int64 `pdefault:"10"`
		Field2 *int64 `pdefault:"20"`
	}

	val1 := int64(30)
	val2 := int64(30)
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

func Test_Int64_InitalisesMultipleFieldWhenSomeAreSet(t *testing.T) {
	type TestStruct struct {
		Field1 *int64 `pdefault:"10"`
		Field2 *int64 `pdefault:"20"`
	}

	val1 := int64(30)
	input := TestStruct{
		Field1: &val1,
	}
	val2 := int64(20)
	expected := TestStruct{
		Field1: &val1,
		Field2: &val2,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_Int64_NonPointerFieldsWithoutPdefaultAreIgnored(t *testing.T) {
	type TestStruct struct {
		Field1 *int64 `pdefault:"10"`
		NonPointerField float64
	}

	input := TestStruct{
		NonPointerField: 1.0,
	}

	val := int64(10)
	expected := TestStruct{
		Field1: &val,
		NonPointerField: 1.0,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_Int64_WhenDefaultNotConvertibleToIntPanicsWhenFieldNil(t *testing.T) {

	// Do not perform this test on strings since a tag will always be convertable
	// to a string type.
	if reflect.TypeOf(int64(10)).Kind() != reflect.String {
		type TestStruct struct {
			Field1 *int64 `pdefault:"someString"`
		}

		input := TestStruct{
		}

		defer func() {
			if panicErr := recover(); panicErr != nil {
				assert.Regexp(t, "(.*someString)(.*int64)(.*Field1)", panicErr, "Error should contain the field name")
			} else {
				assert.Fail(t, "Did not panic")
			}
		}()

		pdefault.Init(&input)
	}
}

func Test_Int64_WhenDefaultNotConvertibleToIntPanicsWhenFieldNotNil(t *testing.T) {

	// Do not perform this test on strings since a tag will always be convertable
	// to a string type.
	if reflect.TypeOf(int64(10)).Kind() != reflect.String {
		type TestStruct struct {
			Field1 *int64 `pdefault:"hello"`
		}

		val := int64(10)
		input := TestStruct{
			Field1: &val,
		}

		defer func() {
			if panicErr := recover(); panicErr != nil {
				assert.Regexp(t, "(.*hello)(.*int64)(.*Field1)", panicErr, "Error should contain the field name")
			} else {
				assert.Fail(t, "Did not panic")
			}
		}()

		pdefault.Init(&input)
	}
}

func Test_Int64_PdefaultTagOnNotPointerFieldPanics(t *testing.T) {

	type TestStruct struct {
		NonPointerField int64 `pdefault:"10"`
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