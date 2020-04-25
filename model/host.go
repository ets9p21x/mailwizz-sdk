package model

const (
	MethodPOST   = "POST"
	MethodGET    = "GET"
	MethodPUT    = "PUT"
	MethodDELETE = "DELETE"
)

type Host struct {
	Method string
	Source string
}
