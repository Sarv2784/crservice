package library

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func StructToStringMap[T any](input T) (map[string]string, error) {
	result := make(map[string]string)

	val := reflect.ValueOf(input)
	typ := reflect.TypeOf(input)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input is not of type struct")
	}

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)

		var strValue string
		switch fieldValue.Kind() {
		case reflect.String:
			strValue = fieldValue.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			strValue = strconv.FormatInt(fieldValue.Int(), 10)
		case reflect.Float32, reflect.Float64:
			strValue = strconv.FormatFloat(fieldValue.Float(), 'f', -1, 64)
		case reflect.Bool:
			strValue = strconv.FormatBool(fieldValue.Bool())
		default:
			strValue = fmt.Sprintf("%v", fieldValue.Interface())
		}
		fieldName := strings.ToLower(field.Name)
		result[strings.ToLower(fieldName)] = strValue
	}

	return result, nil
}
