package utils

import (
	"MailWizz-with-go/model"
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func Signature(h model.Host, formdata string) string {
	wUrl := GobalConfig.Host + h.Source
	hosts := strings.Split(wUrl, "?")
	fmt.Println("hosts: ", hosts)
	if h.Method == model.MethodGET && len(hosts) > 1 {
		pathname, _ := url.Parse(wUrl + "&" + formdata)
		formdata = pathname.Query().Encode()
	} else {
		pathname, _ := url.Parse(wUrl +"?" + formdata)
		formdata = pathname.Query().Encode()
	}
	signature := fmt.Sprintf("%s %s?%s", h.Method, hosts[0], formdata)
	fmt.Println("signature", signature)
	hash := hmac.New(sha1.New, []byte(GobalConfig.PrivateKey))
	hash.Write([]byte(signature))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// StructToMap json to url
func StructToMap(i interface{}) map[string]string {
	var m = make(map[string]string)
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	for j := 0; j < t.NumField(); j++ {
		key := t.Field(j).Name
		val := v.Field(j).Interface()
		switch val.(type) {
		case int:
			m[key] = strconv.Itoa(v.Field(j).Interface().(int))
			break
		default:
			m[key] = v.Field(j).Interface().(string)
			break
		}

	}
	return m
}
