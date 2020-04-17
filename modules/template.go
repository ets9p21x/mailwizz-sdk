package modules

import (
	"MailWizz-with-go/model"
	"MailWizz-with-go/utils"
	"encoding/json"
	"strconv"
)

const templatePath = "/templates"

// GetTemplates 获取template列表
func GetTemplates(page, perPage int) (model.Templates, error) {
	p := model.PageStruct{
		Page:    strconv.Itoa(page),
		PerPage: strconv.Itoa(perPage),
	}
	data := utils.StructToMap(p, "")
	h := model.Host{
		Method: model.MethodGET,
		Source: templatePath,
	}
	r := model.Templates{}
	result, err := utils.HttpGo(h, data)
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(result, &r)
	return r, err
}

// GetTemplate 获取某个列表
func GetTemplate(uid string) (model.Template, error) {
	h := model.Host{
		Method: model.MethodGET,
		Source: templatePath + "/" + uid,
	}
	result, err := utils.HttpGo(h, make(map[string]string))
	r := model.Template{}
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(result, &r)
	return r, err
}
