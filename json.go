package maas

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

/*
maas can change their output json in code easily,
for keeping every thing in obj, we will use our json encoder
decoder to keep everything is in there.
if you don't need you can use code import
*/

var dynamicValue = true

func NoDynamicValue() {
	dynamicValue = false
}

func DynamicValue() {
	dynamicValue = true
}

// Unmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v. If v is nil or not a pointer,
// any data no define in struct, will be set in Obj.X
func Unmarshal(data []byte, result interface{}) error {
	if !dynamicValue {
		return json.Unmarshal(data, result)
	}
	return unmarshal(data, result)
}

func unmarshalSlice(data, result interface{}) error {
	dataVal := reflect.ValueOf(data)
	resultVal := reflect.ValueOf(result)
	resultSliceV := resultVal.Elem()

	elementType := resultVal.Elem().Type().Elem()

	dataLen := dataVal.Len()
	for i := 0; i < dataLen; i++ {
		oneVal := dataVal.Index(i)
		element := reflect.New(elementType)

		b, err := json.Marshal(oneVal.Interface())
		if err != nil {
			return err
		}
		err = unmarshal(b, element.Interface())
		if err != nil {
			return err
		}
		e := element.Interface()
		resultSliceV.Set(reflect.Append(resultSliceV, reflect.ValueOf(e).Elem()))
	}
	return nil
}

func unmarshalMap(data, result interface{}) error {
	temp, ok := data.(map[string]interface{})
	if !ok {
		return &json.UnmarshalTypeError{
			Value: "map[string]interface{}",
			Type:  reflect.TypeOf(data),
		}
	}

	// create a mapstructure decoder
	var md mapstructure.Metadata
	decoder, err := mapstructure.NewDecoder(
		&mapstructure.DecoderConfig{
			TagName:  "json",
			Squash:   true,
			Metadata: &md,
			Result:   result,
		})
	if err != nil {
		return err
	}

	// decode the unmarshalled map into the given struct
	if err := decoder.Decode(temp); err != nil {
		return err
	}

	return nil
}

func unmarshal(data []byte, result interface{}) error {
	rv := reflect.ValueOf(result)
	test := reflect.New(rv.Elem().Type()).Interface()

	// test it can be success unmarshal
	if err := json.Unmarshal(data, test); err != nil {
		return err
	}

	//
	var any interface{}
	if err := json.Unmarshal(data, &any); err != nil {
		return err
	}

	anyVal := reflect.ValueOf(any)
	switch anyVal.Kind() {
	case reflect.Slice:
		return unmarshalSlice(any, result)
	case reflect.Map:
		return unmarshalMap(any, result)
	default:
		return json.Unmarshal(data, result)
	}
}

func Marshal(source interface{}) (data []byte, err error) {
	if !dynamicValue {
		return json.Marshal(source)
	}

	defer func() {
		if e := recover(); e != nil {
			data, err = json.Marshal(source)
		}
	}()

	dataVal := reflect.ValueOf(source)
	dataTyp := dataVal.Type()

	switch dataTyp.Kind() {

	}

	if dataTyp.Kind() != reflect.Struct {
		return json.Marshal(source)
	}

	m := make(map[string]json.RawMessage)

	for i := 0; i < dataVal.NumField(); i++ {
		fieldTyp := dataTyp.Field(i)
		fieldVal := dataVal.Field(i)

		if fieldVal.Type() == ObjType {
			b, e := json.Marshal(fieldVal.Interface())
			if e != nil {
				return nil, e
			}
			var t map[string]json.RawMessage
			e = json.Unmarshal(b, &t)
			if e != nil {
				return nil, e
			}

			for k, v := range t {
				if k == "-" {
					continue
				}
				m[k] = v
			}

			if o, ok := fieldVal.Interface().(Obj); ok {
				for k, v := range o.X {
					m[k], err = json.Marshal(v)
					if err != nil {
						return nil, err
					}
				}
			}

		} else {
			m[fieldTyp.Name], err = Marshal(fieldVal.Interface())
			if err != nil {
				return nil, err
			}
		}
	}

	return json.Marshal(m)
}
