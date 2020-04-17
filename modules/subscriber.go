package modules

import (
	"MailWizz-with-go/model"
	"MailWizz-with-go/utils"
	"encoding/json"
	"strconv"
)

// GetSubscribers 订阅列表
func GetSubscribers(listUid string, page, perPage int) (model.Subscribers, error) {
	p := model.PageStruct{
		Page:    strconv.Itoa(page),
		PerPage: strconv.Itoa(perPage),
	}
	data := utils.StructToMap(p, "")
	h := model.Host{
		Method: model.MethodGET,
		Source: listPath + "/" + listUid + "/subscribers",
	}
	r := model.Subscribers{}
	result, err := utils.HttpGo(h, data)
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(result, &r)
	return r, err
}

// GetSubscriber 订阅列表
func GetSubscriber(listUid, subscriberUid string) (model.Subscriber, error) {

	h := model.Host{
		Method: model.MethodGET,
		Source: listPath + "/" + listUid + "/subscribers/" + subscriberUid,
	}
	r := model.Subscriber{}
	result, err := utils.HttpGo(h, make(map[string]string))
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(result, &r)
	return r, err
}
