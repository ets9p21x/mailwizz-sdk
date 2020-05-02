package modules

import (
	"mailwizz-sdk/model"
	"mailwizz-sdk/param"
	"mailwizz-sdk/utils"
	"encoding/json"
	"strconv"
)

// GetSubscribers get a list of subscribers
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

// GetSubscriber get a subscriber
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

// CreateSubscriber create a subscriber
func CreateSubscriber(listUid string, sub param.Subscriber) (model.SubscriberOne, error) {
	r := model.SubscriberOne{}

	h := model.Host{
		Method: model.MethodPOST,
		Source: listPath + "/" + listUid + "/subscribers",
	}
	data := utils.StructToMap(sub, "")
	result, err := utils.HttpGo(h, data)

	if err != nil {
		return r, err
	}
	err = json.Unmarshal(result, &r)
	return r, err
}

// UpdateSubscriber update the subscriber
func UpdateSubscriber(listUid string, sub param.Subscriber) (model.SubscriberOne, error) {
	r := model.SubscriberOne{}

	h := model.Host{
		Method: model.MethodPUT,
		Source: listPath + "/" + listUid + "/subscribers/" + sub.SubscriberUid,
	}
	data := utils.StructToMap(sub, "")
	result, err := utils.HttpGo(h, data)

	if err != nil {
		return r, err
	}
	err = json.Unmarshal(result, &r)
	return r, err
}

// DeleteSubscriber update the subscriber
func DeleteSubscriber(listUid, subscriberUid string) (model.SubscriberOne, error) {
	r := model.SubscriberOne{}

	h := model.Host{
		Method: model.MethodDELETE,
		Source: listPath + "/" + listUid + "/subscribers/" +subscriberUid,
	}
	result, err := utils.HttpGo(h, nil)

	if err != nil {
		return r, err
	}
	err = json.Unmarshal(result, &r)
	return r, err
}

// SearchSubscriberByEmail search subscriber by email
func SearchSubscriberByEmail(listUid, email string) (model.SubscriberOne, error) {
	r := model.SubscriberOne{}

	h := model.Host{
		Method: model.MethodGET,
		Source: listPath + "/" + listUid + "/subscribers/search-by-email",
	}
	data := map[string]string{
		"EMAIL": email,
	}
	result, err := utils.HttpGo(h, data)

	if err != nil {
		return r, err
	}
	err = json.Unmarshal(result, &r)
	return r, err
}

// Unsubscriber Unsubscribe existing subscriber from the list
func Unsubscriber(listUid, subscriberUid string) (model.SubscriberOne, error) {
	r := model.SubscriberOne{}

	h := model.Host{
		Method: model.MethodPUT,
		Source: listPath + "/" + listUid + "/subscribers/"+ subscriberUid + "/unsubscribe",
	}
	result, err := utils.HttpGo(h, nil)

	if err != nil {
		return r, err
	}
	err = json.Unmarshal(result, &r)
	return r, err
}

// CreateSubscriberInBulk ADD SUBSCRIBERS IN BULK (since MailWizz 1.8.1)
func CreateSubscriberInBulk(listUid string, subs []param.Subscriber) (model.SubscriberOne, error) {
	r := model.SubscriberOne{}

	h := model.Host{
		Method: model.MethodPOST,
		Source: listPath + "/" + listUid + "subscribers/bulk",
	}
	data := make(map[string]string)
	for k, v := range subs {
		t := utils.StructToMap(v, "")
		for ki, vi := range t {
			data["["+strconv.Itoa(k)+"]["+ki+"]"] = vi
		}
	}
	result, err := utils.HttpGo(h, data)

	if err != nil {
		return r, err
	}
	err = json.Unmarshal(result, &r)
	return r, err
}