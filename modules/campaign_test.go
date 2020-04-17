package modules

import (
	"MailWizz-with-go/param"
	"testing"
)

func TestCreateCampaign(t *testing.T) {
	data := param.Campaign{
		Campaign: param.CampaignBody{
			Name:       "dddd",
			Type:       "regular",
			FromName:   "system",
			FromEmail:  "info@xxx.com",
			Subject:    "Happy New Year",
			ReplyTo:    "info@xxx.com",
			SendAt:     "2020-04-20 14:07:00",//time.Now().Format("2006-01-02 15:04:05"),
			ListUid:    "",
			SegmentUid: "",
			Options: param.Options{
				UrlTracking:"no",
				JsonFeed:"no",
				XmlFeed:"no",
				PlainTextEmail:"yes",
			},
			CampaignTemplate: param.CampaignTemplate{
				TemplateUid:"",
				Content:"<!DOCTYPE html><html><head><meta name='charset' content='utf-8'><title></title></head><body>dsdsdssd<br /><br />[UNSUBSCRIBE_URL], [COMPANY_FULL_ADDRESS]</body></html>",
				InlineCss: "no",
				AutoPlainText:"yes",
			},
		},
	}
	r, err := CreateCampaign(data)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(r)
}
