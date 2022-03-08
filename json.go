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
func Unmarshal(data []byte, result interface{}) (err error) {
	if !dynamicValue {
		return json.Unmarshal(data, result)
	}

	// hope no reflect panic here.
	defer func() {
		if e := recover(); e != nil {
			err = json.Unmarshal(data, result)
		}
	}()

	return unmarshal(data, result)
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
	switch {
	case anyVal.Kind() == rv.Elem().Kind() && anyVal.Kind() == reflect.Slice:
		return unmarshalSlice(any, result)
	case anyVal.Kind() == reflect.Map:
		return unmarshalMap(any, result)
	default:
		return json.Unmarshal(data, result)
	}
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
		resultSliceV.Set(reflect.Append(resultSliceV, element.Elem()))
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

func Marshal(source interface{}) (data []byte, err error) {
	if !dynamicValue {
		return json.Marshal(source)
	}

	// hope no reflect panic here.
	defer func() {
		if e := recover(); e != nil {
			data, err = json.Marshal(source)
		}
	}()
	return marshal(source)
}

func marshal(source interface{}) ([]byte, error) {
	dataVal := reflect.ValueOf(source)
	dataTyp := dataVal.Type()

	// test it can be success unmarshal
	if _, err := json.Marshal(source); err != nil {
		return nil, err
	}

	for dataTyp.Kind() == reflect.Ptr {
		dataVal = dataVal.Elem()
		dataTyp = dataVal.Type()
	}

	switch dataTyp.Kind() {
	case reflect.Slice:
		return marshalSlice(source)
	case reflect.Struct:
		return marshalStruct(source)
	default:
		return json.Marshal(source)

	}
}

func marshalSlice(source interface{}) ([]byte, error) {
	sourceVal := reflect.ValueOf(source)

	res := make([]interface{}, 0, sourceVal.Len())
	for i := 0; i < sourceVal.Len(); i++ {
		d, err := marshal(sourceVal.Index(i).Interface())
		if err != nil {
			return nil, err
		}
		var any interface{}
		_ = json.Unmarshal(d, &any)
		res = append(res, any)
	}

	return json.Marshal(res)
}

func marshalStruct(source interface{}) (data []byte, err error) {
	sourceVal := reflect.ValueOf(source)
	sourceTyp := sourceVal.Type()
	m := make(map[string]json.RawMessage)

	for i := 0; i < sourceVal.NumField(); i++ {
		fieldTyp := sourceTyp.Field(i)
		fieldVal := sourceVal.Field(i)

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
			m[fieldTyp.Name], err = marshal(fieldVal.Interface())
			if err != nil {
				return nil, err
			}
		}
	}

	return json.Marshal(m)
}
