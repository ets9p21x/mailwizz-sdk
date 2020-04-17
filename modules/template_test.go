package modules

import "testing"

func TestGetTemplates(t *testing.T) {
	r, err := GetTemplates(1, 10)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(r)
}

func TestGetTemplate(t *testing.T) {
	r, err := GetTemplate("xs147ebaq0ba6")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(r)
}
