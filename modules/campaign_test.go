package modules

import (
	"MailWizz-with-go/param"
	"testing"
	"time"
)

func TestCreateCampaign(t *testing.T) {
	data := param.Campaign{
		Campaign: param.CampaignBody{
			Name:       "新年活动",
			Type:       "regular",
			FromName:   "system",
			FromEmail:  "info@selpte.com",
			Subject:    "Happy New Year",
			ReplyTo:    "info@selpte.com",
			SendAt:     time.Now().Format("2006-01-02 15:04:05"),
			ListUid:    "xp539hntgk917",
			SegmentUid: "",
			Options: param.Options{
				"no",
				"no",
				"no",
				"yes",
				"",
			},
			CampaignTemplate: param.CampaignTemplate{
				"vj315mqc8x5e7",
				"",
				"<!DOCTYPE html><html><head><meta name=\"charset\" content=\"utf-8\"><title></title></head><body>dsdsdssd<br /><br />[UNSUBSCRIBE_URL], [COMPANY_FULL_ADDRESS]</body></html>",
				"no",
				"",
				"yes",
			},
		},
	}
	r, err := CreateCampaign(data)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(r)
}
