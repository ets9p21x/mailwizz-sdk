package param

type Subscriber struct {
	SubscriberUid string `json:"subscriber_uid"`
	Email         string `json:"EMAIL"`
	FName         string `json:"FNAME"`
	LName         string `json:"LNAME"`
	Custom        string `json:"CUSTOM"`
}
