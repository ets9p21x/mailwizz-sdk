package utils

import (
	"MailWizz-with-go/model"
	"testing"
)

func TestHttpGo(t *testing.T) {
	user := model.UserInfo{
		EMAIL:  "1234@qq.com",
		FNAME:  "Will1",
		LNAME:  "Smit1h",
		CUSTOM: "custo1m",
	}
	u := StructToMap(user, "")
	h := model.Host{
		Method: model.MethodPOST,
		Source: "/lists/xp539hntgk917/subscribers",
	}
	_, err := HttpGo(h, u)
	if err != nil {
		t.Error(err)
	}

	h = model.Host{
		Method: model.MethodGET,
		Source: "/lists",
	}

	var g = make(map[string]string)
	g["page"] = "1"
	g["per_page"] = "10"
	_, err = HttpGo(h, g)
	if err != nil {
		t.Error(err)
	}
}
