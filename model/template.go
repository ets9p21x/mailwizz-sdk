package model

type TemplatesRecord struct {
	TemplateUid string `json:"template_uid"`
	Name        string `json:"name"`
	Screenshot  string `json:"screenshot"`
}

type TemplateRecord struct {
	Name       string `json:"name"`
	Content    string `json:"content"`
	InlineCss  string `json:"inline_css"`
	Screenshot string `json:"screenshot"`
}

type TemplatesData struct {
	Count       string            `json:"count"`
	TotalPages  int               `json:"total_pages"`
	CurrentPage int               `json:"current_page"`
	Records     []TemplatesRecord `json:"records"`
}

type Templates struct {
	Status string        `json:"status`
	Data   TemplatesData `json:"data"`
	Error  string        `json:"error"`
}

type TemplateData struct {
	Record TemplateRecord `json:"record"`
}

type Template struct {
	Status string       `json:"status`
	Data   TemplateData `json:"data"`
	Error  string       `json:"error"`
}
