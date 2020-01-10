package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var outputDir = flag.String("output_dir", ".", "Output directory for generated test files")

var testCases = []struct{
	Name string
	Type string
	TagValue1 string // A valid pdefault tag - e.g. 10
	Literal1 string // Literal value equal to TagValue1 - e.g. int64(10)
	TagValue2 string // A valid pdefault tag - e.g. 20
	Literal2 string // Literal value equal to TagValue2 - e.g. int64(20)
	Literal3 string // Literal value not equal to TagValues 1 or 2 - e.g. int64(30)
}{
	{
		Name: "Int64",
		Type: "int64",
		TagValue1: "10",
		Literal1: "int64(10)",
		TagValue2: "20",
		Literal2: "int64(20)",
		Literal3: "int64(30)",
	},
	{
		Name: "Float64",
		Type: "float64",
		TagValue1: "1.1",
		Literal1: "float64(1.1)",
		TagValue2: "2.2",
		Literal2: "float64(2.2)",
		Literal3: "float64(3.3)",
	},
}

func main() {
	flag.Parse()

  t := template.Must(template.New("queue").Parse(testTemplate))

  for _, testCase := range testCases {

		outFileName := filepath.Join(*outputDir, fmt.Sprintf("%s_test.go", strings.ToLower(testCase.Name)))
		outFile, err := os.Create(outFileName)
		if err != nil {
			log.Fatalf("Error opening output file %q: %v", outFileName, err)
		}

		err = t.Execute(outFile, testCase)
		if err != nil {
			panic(errors.Wrapf(err, "Failed to generate test case %q", outFileName))
		}
	}
}

var testTemplate = `package generated_tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/thecodedproject/pdefault"
	"reflect"
	"testing"
)

func Test_{{.Name}}_InitalisesFieldWhenNotSet(t *testing.T) {
	type TestStruct struct {
		Field *{{.Type}} ` + "`pdefault:\"{{.TagValue1}}\"`" + `
	}

	input := TestStruct{}
	val := {{.Literal1}}
	expected := TestStruct{
		Field: &val,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_{{.Name}}_InitalisesMultipleFieldWhenNotSet(t *testing.T) {
	type TestStruct struct {
		Field1 *{{.Type}} ` + "`pdefault:\"{{.TagValue1}}\"`" + `
		Field2 *{{.Type}} ` + "`pdefault:\"{{.TagValue2}}\"`" + `
	}

	input := TestStruct{}
	val1 := {{.Literal1}}
	val2 := {{.Literal2}}
	expected := TestStruct{
		Field1: &val1,
		Field2: &val2,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_{{.Name}}_InitalisesMultipleFieldWhenAllAreSet(t *testing.T) {
	type TestStruct struct {
		Field1 *{{.Type}} ` + "`pdefault:\"{{.TagValue1}}\"`" + `
		Field2 *{{.Type}} ` + "`pdefault:\"{{.TagValue2}}\"`" + `
	}

	val1 := {{.Literal3}}
	val2 := {{.Literal3}}
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

func Test_{{.Name}}_InitalisesMultipleFieldWhenSomeAreSet(t *testing.T) {
	type TestStruct struct {
		Field1 *{{.Type}} ` + "`pdefault:\"{{.TagValue1}}\"`" + `
		Field2 *{{.Type}} ` + "`pdefault:\"{{.TagValue2}}\"`" + `
	}

	val1 := {{.Literal3}}
	input := TestStruct{
		Field1: &val1,
	}
	val2 := {{.Literal2}}
	expected := TestStruct{
		Field1: &val1,
		Field2: &val2,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_{{.Name}}_NonPointerFieldsWithoutPdefaultAreIgnored(t *testing.T) {
	type TestStruct struct {
		Field1 *{{.Type}} ` + "`pdefault:\"{{.TagValue1}}\"`" + `
		NonPointerField float64
	}

	input := TestStruct{
		NonPointerField: 1.0,
	}

	val := {{.Literal1}}
	expected := TestStruct{
		Field1: &val,
		NonPointerField: 1.0,
	}

	pdefault.Init(&input)

	assert.Equal(t, expected, input)
}

func Test_{{.Name}}_WhenDefaultNotConvertibleToIntPanicsWhenFieldNil(t *testing.T) {

	// Do not perform this test on strings since a tag will always be convertable
	// to a string type.
	if reflect.TypeOf({{.Literal1}}).Kind() != reflect.String {
		type TestStruct struct {
			Field1 *{{.Type}} ` + "`pdefault:\"someString\"`" + `
		}

		input := TestStruct{
		}

		defer func() {
			if panicErr := recover(); panicErr != nil {
				assert.Regexp(t, "(.*someString)(.*{{.Type}})(.*Field1)", panicErr, "Error should contain the field name")
			} else {
				assert.Fail(t, "Did not panic")
			}
		}()

		pdefault.Init(&input)
	}
}

func Test_{{.Name}}_WhenDefaultNotConvertibleToIntPanicsWhenFieldNotNil(t *testing.T) {

	// Do not perform this test on strings since a tag will always be convertable
	// to a string type.
	if reflect.TypeOf({{.Literal1}}).Kind() != reflect.String {
		type TestStruct struct {
			Field1 *{{.Type}} ` + "`pdefault:\"hello\"`" + `
		}

		val := {{.Literal1}}
		input := TestStruct{
			Field1: &val,
		}

		defer func() {
			if panicErr := recover(); panicErr != nil {
				assert.Regexp(t, "(.*hello)(.*{{.Type}})(.*Field1)", panicErr, "Error should contain the field name")
			} else {
				assert.Fail(t, "Did not panic")
			}
		}()

		pdefault.Init(&input)
	}
}

func Test_{{.Name}}_PdefaultTagOnNotPointerFieldPanics(t *testing.T) {

	type TestStruct struct {
		NonPointerField {{.Type}} ` + "`pdefault:\"{{.TagValue1}}\"`" + `
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
`
