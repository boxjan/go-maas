package maas

import (
	"testing"
	"time"
)

func TestObj_AssignClient(t *testing.T) {

	type BooObj struct {
		Obj
		Time time.Time `json:"time"`
	}

	type FooObj struct {
		Obj
		BooObj BooObj `json:"boo_obj"`
	}

	c := Client{}

	f := FooObj{Obj: Obj{client: &c}}

	if f.BooObj.client != &c {
		t.Fail()
	}
}
