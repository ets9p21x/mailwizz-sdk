package modules

import (
	"MailWizz-with-go/param"
	"testing"
)

func TestGetSubscribers(t *testing.T) {
	r, err := GetSubscribers("xp539hntgk917", 1, 10)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(r)
}

func TestGetSubscriber(t *testing.T) {
	r, err := GetSubscriber("xp539hntgk917", "ml5150wh38x19b")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(r)
}

func TestCreateSubscriber(t *testing.T) {
	data := param.Subscriber{
		Email:  "admin1@outlook.com",
		FName:  "will1",
		LName:  "smith1",
		Custom: "custom1",
	}
	r, err := CreateSubscriber("xp539hntgk917", data)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(r)
}

func TestUpdateSubscriber(t *testing.T) {
	data := param.Subscriber{
		SubscriberUid: "zx160tgze7892",
		Email:  "admin@outlook.com",
		FName:  "will1",
		LName:  "smith2",
		Custom: "custom",
	}
	r, err := UpdateSubscriber("xp539hntgk917", data)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(r)
}

func TestDeleteSubscriber(t *testing.T) {
	r, err := DeleteSubscriber("xp539hntgk917", "zx160tgze7892")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(r)
}

// not pass
func TestSearchSubscriberByEmail(t *testing.T) {
	r, err := SearchSubscriberByEmail("xp539hntgk917", "admin@outlook.com")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(r)
}

func TestUnsubscriber(t *testing.T) {
	r, err := Unsubscriber("xp539hntgk917", "ax4755kkhd0a1")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(r)
}

