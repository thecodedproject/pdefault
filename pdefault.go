package pdefault

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strconv"

	"log"
)

var conversions = map[reflect.Kind]func(string) (reflect.Value, error){
	reflect.Float32: func(v string) (reflect.Value, error) {
		val64, err := strconv.ParseFloat(v, 32)
		val := float32(val64)

		log.Println("val:", val)

		return reflect.ValueOf(&val), err
	},
	reflect.Float64: func(v string) (reflect.Value, error) {
		val, err := strconv.ParseFloat(v, 64)
		return reflect.ValueOf(&val), err
	},
	reflect.Int: func(v string) (reflect.Value, error) {
		val64, err := strconv.ParseInt(v, 10, 64)
		// TODO check for overflow in case of running on 32-bit system
		val := int(val64)
		return reflect.ValueOf(&val), err
	},
	reflect.Int8: func(v string) (reflect.Value, error) {
		val64, err := strconv.ParseInt(v, 10, 64)
		// TODO check for overflow in case of running on 32-bit system
		val := int8(val64)
		return reflect.ValueOf(&val), err
	},
	reflect.Int16: func(v string) (reflect.Value, error) {
		val64, err := strconv.ParseInt(v, 10, 64)
		// TODO check for overflow in case of running on 32-bit system
		val := int16(val64)
		return reflect.ValueOf(&val), err
	},
	reflect.Int32: func(v string) (reflect.Value, error) {
		val64, err := strconv.ParseInt(v, 10, 64)
		// TODO check for overflow in case of running on 32-bit system
		val := int32(val64)
		return reflect.ValueOf(&val), err
	},
	reflect.Int64: func(v string) (reflect.Value, error) {
		val, err := strconv.ParseInt(v, 10, 64)
		return reflect.ValueOf(&val), err
	},
	reflect.String: func(v string) (reflect.Value, error) {
		return reflect.ValueOf(&v), nil
	},
	reflect.Uint: func(v string) (reflect.Value, error) {
		val64, err := strconv.ParseUint(v, 10, 64)
		// TODO check for overflow in case of running on 32-bit system
		val := uint(val64)
		return reflect.ValueOf(&val), err
	},
	reflect.Uint8: func(v string) (reflect.Value, error) {
		val64, err := strconv.ParseUint(v, 10, 8)
		val := uint8(val64)
		return reflect.ValueOf(&val), err
	},
	reflect.Uint16: func(v string) (reflect.Value, error) {
		val64, err := strconv.ParseUint(v, 10, 16)
		val := uint16(val64)
		return reflect.ValueOf(&val), err
	},
	reflect.Uint32: func(v string) (reflect.Value, error) {
		val64, err := strconv.ParseUint(v, 10, 32)
		val := uint32(val64)
		return reflect.ValueOf(&val), err
	},
	reflect.Uint64: func(v string) (reflect.Value, error) {
		val, err := strconv.ParseUint(v, 10, 64)
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
