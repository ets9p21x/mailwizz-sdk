package utils

import (
	"mailwizz-sdk/model"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// HttpGo 发送数据到email服务器
func HttpGo(h model.Host, data map[string]string) ([]byte, error) {
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
	log.Println(formdata.Encode())
	hash := Signature(h, mts.Encode())
	req, _ := http.NewRequest(h.Method, GobalConfig.Host+h.Source, bytes.NewBuffer([]byte(formdata.Encode())))
	req.Header.Add("X-MW-PUBLIC-KEY", GobalConfig.PublicKey)
	req.Header.Add("X-MW-TIMESTAMP", timestamp)
	req.Header.Add("X-MW-REMOTE-ADDR", "")
	req.Header.Add("X-MW-SIGNATURE", hash)
	req.Header.Add("X-HTTP-Method-Override", h.Method)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(formdata.Encode())))

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
