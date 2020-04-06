package utils

import (
	"MailWizz-with-go/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// HttpGo 发送数据到email服务器
func HttpGo(h model.Host, data map[string]string) (interface{}, string){
	timestamp := strconv.Itoa(int(time.Now().Unix()))

	formdata := url.Values{}
	mts := url.Values{}
	for k, v := range data {
		formdata.Set(k, v)
		mts.Set(k, v)
	}

	mts.Set("X-MW-PUBLIC-KEY", GobalConfig.PublicKey)
	mts.Set("X-MW-TIMESTAMP", timestamp)
	mts.Set("X-MW-REMOTE-ADDR", "")

	hash := Signature(h, mts.Encode())


	req, _ := http.NewRequest(h.Method, GobalConfig.Host + h.Source, strings.NewReader(formdata.Encode()))
	req.Header.Add("X-MW-PUBLIC-KEY", GobalConfig.PublicKey)
	req.Header.Add("X-MW-TIMESTAMP", timestamp)
	req.Header.Add("X-MW-REMOTE-ADDR", "")
	req.Header.Add("X-MW-SIGNATURE", hash)
	req.Header.Add("X-HTTP-Method-Override", h.Method)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(formdata.Encode())))

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, "Post failed: " + h.Source + "-" + err.Error()
	}

	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err.Error()
	}
	fmt.Println("result: ", string(result))
	return result, ""
}