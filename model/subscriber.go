package model

type SubscribersRecord struct {
	SubscriberUid string `json:"subscriber_uid"`
	EMAIL         string `json:"EMAIL"`
	FNAME         string `json:"FNAME"`
	LNAME         string `json:"LNAME"`
	Status        string `json:"status"`
	Source        string `json:"source"`
	IpAddress     string `json:"ip_address"`
}

type SubscribersData struct {
	Count       string              `json:"count"`
	TotalPages  int                 `json:"total_pages"`
	CurrentPage int                 `json:"current_page"`
	Records     []SubscribersRecord `json:"records"`
}

type Subscribers struct {
	Status string          `json:"status`
	Data   SubscribersData `json:"data"`
	Error  string          `json:"error"`
}

type SubscriberData struct {
	Record SubscribersRecord `json:"record"`
}

type Subscriber struct {
	Status string         `json:"status`
	Data   SubscriberData `json:"data"`
	Error  string         `json:"error"`
}

type SubscriberOneDataRecord struct {
	SubscriberUid string `json:"subscriber_uid"`
	Email         string `json:"email"`
	IpAddress     string `json:"ip_address"`
	Source        string `json:"source"`
}

type SubscriberOneData struct {
	Record SubscriberOneDataRecord `json:"record"`
}

type SubscriberOne struct {
	Status string            `json:"status`
	Data   SubscriberOneData `json:"data"`
	Error  string            `json:"error"`
}
