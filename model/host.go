package model

const (
	MethodPOST = "POST"
	MethodGET  = "GET"
)

type Host struct {
	Method string
	Source string
}
