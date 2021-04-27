package util

import "reflect"

func StructToMap(item interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	if item == nil {
		return result
	}

	v := reflect.TypeOf(item)
	reflectValue := reflect.Indirect(reflect.ValueOf(item))

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := reflectValue.Type().Field(i).Name
		value := reflectValue.Field(i).Interface()
		result[field] = value
	}

	return result
}
