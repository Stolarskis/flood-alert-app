package handler

import (
	"fmt"
	"testing"
)

func TestAddUserData(t *testing.T) {
	testUser := &user{
		Name:  "testname",
		Email: "testemail@tealkgrae",
		Phone: "23454535663",
	}

	err := testUser.addUserData()
	if err != nil {
		fmt.Println(err.Error())
	}

}
