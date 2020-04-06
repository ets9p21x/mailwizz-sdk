package utils

import (
	"MailWizz-with-go/model"
	"fmt"
	"testing"
)

func TestSignature(t *testing.T) {
	signatures := [3]string{
		"eb2e4cccfa0b77a91da7cbbccf8d18ed980d1c50",
		"32c16c8de060c76c07774dbe7701edf4e552335e",
		"35a39acbf298c01f841592651f6af39d83d69d42",
	}
	m := [3]model.Host{
		model.Host{
			Method: "POST",
			Source: "/mail",
		},
		model.Host{
			Method: "GET",
			Source: "/mail",
		},
		model.Host{
			Method: "GET",
			Source: "/mail?d=1",
		},
	}
	for k, v := range m {
		if signatures[k] != Signature(v, "b=33@qq.com&c=3&a=12") {
			t.Error("failed")
		}
	}

}

type Std struct {
	Name string
	Age  int
}

func TestStructToMap(t *testing.T) {
	s := Std{
		"will",
		18,
	}

	r := StructToMap(s)
	fmt.Println(r)
}

//func TestMapToQueryString(t *testing.T) {
//	m := map[string]interface{}{
//		"Home": "1",
//		"a": "222",
//		"X-Hss-ss": "2213",
//		"X-Bss-ss": "222",
//		"AXCC": "w2@qq.com",
//	}
//	r := MapToQueryString(m)
//	fmt.Println(r)
//}
