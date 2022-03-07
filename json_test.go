package maas

import (
	"github.com/smartystreets/assertions"
	"testing"
)

func TestUnmarshalJson(t *testing.T) {
	type simpleDomain struct {
		Obj
		Authoritative bool        `json:"authoritative"`
		Ttl           interface{} `json:"ttl"`
		Id            int         `json:"id"`
	}

	sourceDomain := Domain{
		Obj:                 Obj{ResourceUri: "aaaa", X: map[string]interface{}{"1": 2}},
		Authoritative:       true,
		Ttl:                 60,
		Id:                  0,
		ResourceRecordCount: 45,
		IsDefault:           true,
		Name:                "just-for-test",
	}

	tt := simpleDomain{}
	data, err := Marshal(sourceDomain)
	if err != nil {
		t.Fatal(err)
	}

	if err := Unmarshal(data, &tt); err != nil {
		t.Fatal(err)
	}

	assertions.ShouldEqual(tt, sourceDomain)

}
