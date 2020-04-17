package modules

import "testing"

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
