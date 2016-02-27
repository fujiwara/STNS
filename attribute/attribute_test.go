package attribute

import (
	"reflect"
	"testing"
)

func TestGetByName(t *testing.T) {
	users := UserGroups{
		"test1": &All{
			Id:       1,
			LinkUsers: []string{"foo", "bar"},
		},
	}
	_users := users.GetByName("test1")

	for n, u := range _users {
		if u.Id != 1 {
			t.Error("ummatch user id")
		}

		if !reflect.DeepEqual(u.LinkUsers, []string{"foo", "bar"}) {
			t.Error("ummatch link user")
		}

		if n != "test1" {
			t.Error("ummatch user name")
		}
	}
	notfound := users.GetByName("test2")
	if notfound != nil {
		t.Error("ummatch user id")
	}
}
func TestGetById(t *testing.T) {
	users := UserGroups{
		"test1": &All{
			Id: 1,
		},
	}
	_users := users.GetById("1")
	for n, u := range _users {
		if u.Id != 1 {
			t.Error("ummatch user id")
		}

		if n != "test1" {
			t.Error("ummatch user name")
		}
	}

	notfound := users.GetByName("test2")
	if notfound != nil {
		t.Error("ummatch user id")
	}
}