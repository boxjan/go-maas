package maas

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

/*
maas can change their output json in code easily,
for keeping every thing in obj, we will use our json
decoder to keep everything is in there.
*/

func Unmarshal(data []byte, result interface{}) error {
	temp := make(map[string]interface{})
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
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
	defer func() {
		if e := recover(); e != nil {
			data, err = json.Marshal(source)
		}
	}()

	dataVal := reflect.ValueOf(source)
	dataTyp := dataVal.Type()

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
