package model

type SMSBoom struct {
	Desc   string            `json:"desc"`
	Url    string            `json:"url"`
	Method string            `json:"method"`
	Header map[string]string `json:"header"`
	Data   any               `json:"data"`
}

func NewSMSBoom(
	desc string,
	url string,
	method string,
	header map[string]string,
	data map[string]any) *SMSBoom {
	return &SMSBoom{Desc: desc, Url: url, Method: method, Header: header, Data: data}
}
