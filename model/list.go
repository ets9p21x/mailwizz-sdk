package model

// UserInfo 用户信息测试结构
type UserInfo struct {
	EMAIL  string `json:"EMAIL"`
	FNAME  string `json:"FNAME"`
	LNAME  string `json:"LNAME"`
	CUSTOM string `json:"CUSTOM"`
}

// PageStruct
type PageStruct struct {
	Page    string `json:"PAGE"`
	PerPage string `json:"PER_PAGE"`
}

type ListsNotifications struct {
	Subscribe     string `json:"subscribe"`
	Unsubscribe   string `json:"unsubscribe"`
	SubscribeTo   string `json:"subscribe_to"`
	UnsubscribeTo string `json:"unsubscribe_to"`
}

type ListsCountry struct {
	CountryId string `json:"country_id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
}

type ListsCompany struct {
	Name          string       `json:"name"`
	Address1      string       `json:"address_1"`
	Address2      string       `json:"address_2"`
	ZoneName      string       `json:"zone_name"`
	City          string       `json:"city"`
	ZipCode       string       `json:"zip_code"`
	Phone         string       `json:"phone"`
	AddressFormat string       `json:"address_format"`
	Country       ListsCountry `json:"country"`
}

type ListsGeneral struct {
	ListsUid    string `json:"list_uid"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
}

type ListsDefaults struct {
	FromName string `json:"from_name"`
	ReplyTo  string `json:"reply_to"`
	Subject  string `json:"subject"`
}

type ListsRecord struct {
	General       ListsGeneral       `json:"general"`
	Defaults      ListsDefaults      `json:"defaults"`
	Notifications ListsNotifications `json:"notifications"`
	Company       ListsCompany       `json:"company"`
}

type ListsData struct {
	Count       string        `json:"count"`
	TotalPages  int           `json:"total_pages"`
	CurrentPage int           `json:"current_page"`
	Records     []ListsRecord `json:"records"`
}

//Lists 列表结构
type Lists struct {
	Status string    `json:"status"`
	Data   ListsData `json:"data"`
	Error  string    `json:"error"`
}

type ListData struct {
	Record ListsRecord `json:"record"`
}

// List List结构
type List struct {
	Status string   `json:"status"`
	Data   ListData `json:"data"`
	Error  string   `json:"error"`
}
