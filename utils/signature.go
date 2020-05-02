package utils

import (
	"mailwizz-sdk/model"
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
	if h.Method == model.MethodGET && len(hosts) > 1 {
		pathname, _ := url.Parse(wUrl + "&" + formdata)
		formdata = pathname.Query().Encode()
	} else {
		pathname, _ := url.Parse(wUrl + "?" + formdata)
		formdata = pathname.Query().Encode()
	}
	signature := fmt.Sprintf("%s %s?%s", h.Method, hosts[0], formdata)
	fmt.Println("signature: ", signature)
	hash := hmac.New(sha1.New, []byte(GobalConfig.PrivateKey))
	hash.Write([]byte(signature))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// StructToMap json to url
func StructToMap(i interface{}, s string) map[string]string {
	var m = make(map[string]string)
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	for j := 0; j < t.NumField(); j++ {
		key := t.Field(j).Tag.Get("json")
		val := v.Field(j).Interface()
		switch fmt.Sprintf("%T", val) {
		case "int":
			if len(s) == 0 {
				m[key] = strconv.Itoa(v.Field(j).Interface().(int))
			} else {
				m[s+"["+key+"]"] = strconv.Itoa(v.Field(j).Interface().(int))
			}
			break
		case "string":
			if len(s) == 0 {
				m[key] = v.Field(j).Interface().(string)
			} else {
				m[s+"["+key+"]"] = v.Field(j).Interface().(string)
			}
			break
		default:
			if len(s) == 0 {
				n := StructToMap(val, key)
				m = MapMerge(n, m)
			} else {
				n := StructToMap(val, s+"["+key+"]")
				m = MapMerge(n, m)
			}

			break
		}

	}
	return m
}

func MapMerge(src, dist map[string]string) map[string]string {
	for k, v := range src {
		dist[k] = v
	}
	return dist
}
