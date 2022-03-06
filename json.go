package maas

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

/*
maas can change their output json in code easily,
for keeping every thing in obj, we will use our json encoder
and decoder to keep everything is in there.
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

func Marshal(source interface{}) ([]byte, error) {
	if data, err := marshal(source); err != nil {
		return nil, err
	} else {
		return json.Marshal(data)
	}
}

func marshal(source interface{}) (data map[string]interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			var d2 map[string]interface{}
			if d1, e1 := json.Marshal(source); e1 != nil {
				data = nil
				err = e1
			} else if e2 := json.Unmarshal(d1, &d2); e2 != nil {
				data = nil
				err = e2
			} else {
				data = d2
				err = nil
			}
		}
	}()

}
