package modules

import (
	"MailWizz-with-go/model"
	"MailWizz-with-go/utils"
	"encoding/json"
	"strconv"
)

const listPath = "/lists"

// GetLists 获取list列表
func GetLists(page, perPage int) (model.Lists, error) {
	p := model.PageStruct{
		Page:    strconv.Itoa(page),
		PerPage: strconv.Itoa(perPage),
	}
	data := utils.StructToMap(p, "")
	h := model.Host{
		Method: model.MethodGET,
		Source: listPath,
	}
	r := model.Lists{}
	result, err := utils.HttpGo(h, data)
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(result, &r)
	return r, err
}

// GetList 获取某个列表
func GetList(uid string) (model.List, error) {
	h := model.Host{
		Method: model.MethodGET,
		Source: listPath + "/" + uid,
	}
	result, err := utils.HttpGo(h, make(map[string]string))
	r := model.List{}
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(result, &r)
	return r, err
}
