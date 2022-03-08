package maas

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type simpleSubObj struct {
	Obj
}

var users = []User{
	{
		Obj:         Obj{ResourceUri: "/user/a"},
		IsSuperUser: false,
		IsLocal:     false,
		UserName:    "a",
		Email:       "a@example.com",
	},
	{
		Obj:         Obj{ResourceUri: "/user/b"},
		IsSuperUser: true,
		IsLocal:     false,
		UserName:    "b",
		Email:       "b@example.com",
	},
	{
		Obj:         Obj{ResourceUri: "/user/c"},
		IsSuperUser: false,
		IsLocal:     true,
		UserName:    "c",
		Email:       "c@example.com",
	},
}

var xUsers = []simpleSubObj{
	{
		Obj: Obj{
			X: map[string]interface{}{
				"is_superuser": false,
				"is_local":     false,
				"username":     "a",
				"email":        "a@example.com",
			},
			ResourceUri: "/user/a"},
	},
	{
		Obj: Obj{
			X: map[string]interface{}{
				"is_superuser": true,
				"is_local":     false,
				"username":     "b",
				"email":        "b@example.com",
			},
			ResourceUri: "/user/b"},
	},
	{
		Obj: Obj{
			X: map[string]interface{}{
				"is_superuser": false,
				"is_local":     true,
				"username":     "c",
				"email":        "c@example.com",
			},
			ResourceUri: "/user/c"},
	},
}

var userJson = `{"resource_uri":"/user/a","is_superuser":false,"is_local":false,"username":"a","email":"a@example.com"}`
var usersJson = `[{"resource_uri":"/user/a","is_superuser":false,"is_local":false,"username":"a","email":"a@example.com"},{"resource_uri":"/user/b","is_superuser":true,"is_local":false,"username":"b","email":"b@example.com"},{"resource_uri":"/user/c","is_superuser":false,"is_local":true,"username":"c","email":"c@example.com"}]`

func TestUnmarshal(t *testing.T) {
	Convey("test unmarshal one user into User struct", t, func() {
		u := User{}
		err := Unmarshal([]byte(userJson), &u)

		Convey("should not have any err", func() {
			So(err, ShouldBeNil)
		})
		Convey("result should eq", func() {
			So(u, ShouldResemble, users[0])
		})
	})

	Convey("test unmarshal one user into interface", t, func() {
		var u interface{}
		err := Unmarshal([]byte(userJson), &u)

		Convey("should not have any err", func() {
			So(err, ShouldBeNil)
		})
	})

	Convey("test unmarshal all user into User struct", t, func() {
		var u []User
		err := Unmarshal([]byte(usersJson), &u)

		Convey("should not have any err", func() {
			So(err, ShouldBeNil)
		})
		Convey("result should eq", func() {
			So(u, ShouldResemble, users)
		})
	})

	Convey("test unmarshal all user into interface", t, func() {
		var u interface{}
		err := Unmarshal([]byte(usersJson), &u)

		Convey("should not have any err", func() {
			So(err, ShouldBeNil)
		})
	})

	Convey("test unmarshal one user into simpleSubObj struct", t, func() {
		u := simpleSubObj{}
		err := Unmarshal([]byte(userJson), &u)

		Convey("should not have any err", func() {
			So(err, ShouldBeNil)
		})
		Convey("result should eq", func() {
			So(u, ShouldResemble, xUsers[0])
		})
	})

	Convey("test unmarshal all user into simpleSubObj struct", t, func() {
		var u []simpleSubObj
		err := Unmarshal([]byte(usersJson), &u)

		Convey("should not have any err", func() {
			So(err, ShouldBeNil)
		})
		Convey("result should eq", func() {
			So(u, ShouldResemble, xUsers)
		})
	})
}

func TestMarshal(t *testing.T) {
	Convey("test marshal one user", t, func() {
		b, err := Marshal(xUsers[0])
		Convey("should not have any err", func() {
			So(err, ShouldBeNil)
		})
		Convey("result should eq", func() {
			So(string(b), ShouldEqualJSON, userJson)
		})
	})

	Convey("test marshal all user", t, func() {
		b, err := Marshal(xUsers)
		Convey("should not have any err", func() {
			So(err, ShouldBeNil)
		})
		Convey("result should eq", func() {
			So(string(b), ShouldEqualJSON, usersJson)
		})
	})
}

func TestJsonMarshalAndUnmarshal(t *testing.T) {
	Convey("marshal one user from xUser, then unmarshal to User", t, func() {
		b, err := Marshal(xUsers[0])
		So(err, ShouldBeNil)

		var u User
		err = Unmarshal(b, &u)
		So(err, ShouldBeNil)

		Convey("result should eq", func() {
			So(u, ShouldResemble, users[0])
		})
	})

	Convey("marshal all users from xUser, then unmarshal to Users", t, func() {
		b, err := Marshal(xUsers)
		So(err, ShouldBeNil)

		var u []User
		err = Unmarshal(b, &u)
		So(err, ShouldBeNil)

		Convey("result should eq", func() {
			So(u, ShouldResemble, users)
		})
	})

	Convey("marshal one users from xUser, then unmarshal to User", t, func() {
		b, err := Marshal(xUsers[0])
		So(err, ShouldBeNil)

		var u []User
		err = Unmarshal(b, &u)
		So(err, ShouldNotBeNil)
	})

	Convey("disable dynamicValue", t, func() {
		NoDynamicValue()
		defer DynamicValue()
		b, err := json.Marshal(users)
		So(err, ShouldBeNil)

		u, err := Marshal(users)
		So(err, ShouldBeNil)

		Convey("result should eq", func() {
			So(string(b), ShouldEqualJSON, string(u))
		})
	})

	Convey("disable dynamicValue", t, func() {
		NoDynamicValue()
		defer DynamicValue()

		var anyU, anyT interface{}

		err := json.Unmarshal([]byte(usersJson), &anyU)
		So(err, ShouldBeNil)

		err = Unmarshal([]byte(usersJson), &anyT)
		So(err, ShouldBeNil)

		Convey("result should eq", func() {
			So(anyT, ShouldResemble, anyU)
		})
	})

	Convey("marshal all users from xUser, then unmarshal to User", t, func() {
		b, err := Marshal(xUsers)
		So(err, ShouldBeNil)

		var u User
		err = Unmarshal(b, &u)
		So(err, ShouldNotBeNil)
	})

}
