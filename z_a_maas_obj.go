package maas

import (
	"reflect"
)

type Obj struct {
	ResourceUri string                 `json:"resource_uri"`
	X           map[string]interface{} `json:"-,remain"`
}

type ObjInterface interface {
	getClient() *Client
	setClient(client *Client)
	recursiveClient()
}

var ObjInterfaceType = reflect.TypeOf((*ObjInterface)(nil)).Elem()
var ObjType = reflect.TypeOf(Obj{})

//func recursiveClient(ptr interface{}, p reflect.Type) {
//	for i := p.NumField() - 1; i >= 0; i-- {
//		if p.Field(i).Type.Implements(ObjInterfaceType) {
//			value := reflect.ValueOf(ptr).Elem().Field(i)
//			if !value.IsNil() {
//				//value.SetPointer(unsafe.Pointer(o.getClient()))
//				value.MethodByName("AssignClient").Call([]reflect.Value{})
//			}
//		}
//	}
//}
//
//func objAssignClient(o ObjInterface) {
//	oRef := reflect.TypeOf((**ObjInterface)(unsafe.Pointer(&o)))
//
//	for i := oRef.NumField() - 1; i >= 0; i-- {
//		fmt.Println(oRef.Field(i).Name, oRef.Field(i).Type)
//
//		if oRef.Field(i).Type.Implements(ObjInterfaceType) {
//			value := reflect.ValueOf(o).Elem().Field(i)
//			if !value.IsNil() {
//				//value.SetPointer(unsafe.Pointer(o.getClient()))
//				value.MethodByName("AssignClient").Call([]reflect.Value{})
//			}
//		}
//	}
//}
