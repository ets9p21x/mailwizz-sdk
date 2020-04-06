package main

import (
	"MailWizz-with-go/model"
	"MailWizz-with-go/utils"
	"encoding/json"
)

func init() {
	utils.ReadYML("config/config.yaml")
}

func main() {
	user := model.UserInfo{
		EMAIL:  "346445624@qq.com",
		FNAME:  "Will",
		LNAME:  "Smith",
		CUSTOM: "",
	}
	data, _ := json.Marshal(user)
	utils.Post("/lists/xp539hntgk917/subscribers", data)
}
