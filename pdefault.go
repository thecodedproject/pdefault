package pdefault

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strconv"
)

var conversions = map[reflect.Kind]func(string) (reflect.Value, error){
	reflect.Int64: func(v string) (reflect.Value, error) {
		val, err := strconv.ParseInt(v, 10, 64)
		return reflect.ValueOf(&val), err
	},
	reflect.Float64: func(v string) (reflect.Value, error) {
		val, err := strconv.ParseFloat(v, 64)
		return reflect.ValueOf(&val), err
	},
}

func Init(s interface{}) {

	sValue := reflect.ValueOf(s).Elem()
	sType := sValue.Type()

	for i := 0; i < sType.NumField(); i++ {

		fieldType := sType.Field(i)

		tagValue, tagFound := fieldType.Tag.Lookup("pdefault")
		if !tagFound {
			continue
		}

		if fieldType.Type.Kind() != reflect.Ptr {
			panic(errors.New(
				fmt.Sprintf("Non-pointer field %q with pdefault tag", sType.Field(i).Name),
			))
		}

		val, err := conversions[fieldType.Type.Elem().Kind()](tagValue)
		if err != nil {
			panic(errors.Wrapf(err,
				"Error converting %q to %s for field %s",
				tagValue,
				fieldType.Type.Elem().Kind(),
				fieldType.Name))
		}

		fieldValue := sValue.Field(i)
		if fieldValue.IsNil() {
			fieldValue.Set(val)
		}

	}

}
