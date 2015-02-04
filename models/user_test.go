package models

import "testing"
import "github.com/nemowen/htime/utils"

func TestAddUser(t *testing.T) {
	user := &User{
		Username: "nemowen",
		Password: utils.EncodeByMd5("wenbin"),
		Email:    "wenbin171@163.com",
	}
	err := user.SaveUser()
	if err != nil {
		t.Error("insert user error", err.Error())
	}
}
