package maas

type UndefinedType interface{}

type UndefinedStruct map[string]UndefinedType

func (u UndefinedStruct) Get(k string) interface{} {
	return u[k]
}

func (u *UndefinedStruct) setClient(client *Client) {

}

func (u *UndefinedStruct) getClient() *Client {
	return nil
}

func (u *UndefinedStruct) recursiveClient() {

}
