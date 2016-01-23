package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

type Ambassador struct {
	Ambassador    string
	Path          string
	HTTPVerb      string
	Configuration json.RawMessage
}

var HTTPVerbFunctionMap = map[string]interface{}{
	"GET":  http.Get,
	"HEAD": http.Head,
	"POST": http.Post,
}

// http://stackoverflow.com/a/26746461
func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

func ParseConfiguration(jsonData json.RawMessage, obj interface{}) interface{} {
	fields := map[string]string{}
	err := json.Unmarshal([]byte(jsonData), &fields)
	if err != nil {
		panic(err)
	}
	for k, v := range fields {
		err := SetField(obj, k, v)
		if err != nil {
			panic(err)
		}
	}
	return obj
}
