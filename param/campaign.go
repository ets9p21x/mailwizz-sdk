package param

// Options optional block, defaults are shown
type Options struct {
	UrlTracking    string `json:"url_tracking"`     // yes | no default: no
	JsonFeed       string `json:"json_feed"`        // yes | no default: no
	XmlFeed        string `json:"xml_feed"`         // yes | no default: no
	PlainTextEmail string `json:"plain_text_email"` // yes | no default: yes
	EmailStats     string `json:"email_stats"`      // a valid email address where we should send the stats after campaign done
}

// CampaignTemplate required block, archive or template_uid or content => required.
type CampaignTemplate struct {
	TemplateUid   string `json:"template_uid"`    // required
	Archive       string `json:"archive"`         //
	Content       string `json:"content"`         //
	InlineCss     string `json:"inline_css"`      // yes | no default: no
	PlainText     string `json:"plain_text"`      // leave empty to auto generate
	AutoPlainText string `json:"auto_plain_text"` // yes | no default: yes
}

type CampaignBody struct {
	Name            string          `json:"name"`        // required
	Type            string          `json:"type"`        // optional: regular or autoresponder
	FromName        string          `json:"from_name"`   // required
	FromEmail       string          `json:"from_email"`  // required
	Subject         string          `json:"subject"`     // required
	ReplyTo         string          `json:"reply_to"`    // required
	SendAt          string          `json:"send_at"`     // required, this will use the timezone which customer selected. YYYY-MM-DD hh:mm:ss
	ListUid         string          `json:"list_uid"`    // required
	SegmentUid      string          `json:"segment_uid"` // optional, only to narrow down
	Options         Options         `json:"options"`     //
	CampaignTemplate CampaignTemplate `json:"template"`    //
}

type Campaign struct {
	Campaign CampaignBody `json:"campaign"`
}
