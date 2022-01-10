package maas

type UndefinedType interface{}

type UndefinedStruct map[string]UndefinedType

func (u UndefinedStruct) Get(k string) interface{} {
	return u[k]
}

type T struct {
}
