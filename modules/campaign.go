package modules

import (
	"mailwizz-sdk/model"
	"mailwizz-sdk/param"
	"mailwizz-sdk/utils"
	"encoding/json"
)

const CampaignPath = "/campaigns"

// CreateCampign 创建活动
func CreateCampaign(c param.Campaign) (model.Campaign, error) {
	h := model.Host{
		Method: model.MethodPOST,
		Source: CampaignPath,
	}
	data := utils.StructToMap(c, "")
	r := model.Campaign{}
	result, err := utils.HttpGo(h, data)
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(result, &r)
	return r, err
}
