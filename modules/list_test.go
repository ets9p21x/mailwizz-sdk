package modules

import (
	"fmt"
	"testing"
)

func TestGetLists(t *testing.T) {
	r, err := GetLists(1, 10)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("TestGetLists: ", r)
}

func TestGetList(t *testing.T) {
	r, err := GetList("np9523np1fde2")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("TestGetList: ", r)
}
