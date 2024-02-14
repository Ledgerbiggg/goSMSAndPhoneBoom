package model

import (
	"goSMSBoom/utils"
)

// PhoneMessage 结构体包含了给定的参数
type PhoneMessage struct {
	Seq         int    `json:"seq"`
	MessageID   string `json:"message_id"`
	MsgID       string `json:"msgid"`
	MessageTime int64  `json:"message_time"`
	ImID        string `json:"imid"`
	Paid        int    `json:"paid"`
	ImType      int    `json:"im_type"`
	Source      string `json:"source"`
	SSID        string `json:"ssid"`
	ContentType string `json:"content_type"`
	Content     string `json:"content"`
	Trigger     string `json:"trigger"`
	FromID      string `json:"fromid"`
	SourceTag   string `json:"source_tag"`
	ShowType    string `json:"showType"`
	UserID      int    `json:"userid"`
	PVKey       string `json:"pvKey"`
	WSID        string `json:"wsid"`
	Plat        string `json:"plat"`
	XST         string `json:"xst"`
	AppID       string `json:"appid"`
	Feature     string `json:"feature"`
}

// NewPhoneMessage 是 PhoneMessage 的构造函数，传入 Content 返回一个 PhoneMessage 实例
func NewPhoneMessage(content string, IMID string, SSID string) *PhoneMessage {
	return &PhoneMessage{
		Seq:         1,
		MessageID:   "key" + "18da71384555cd",
		MsgID:       "key" + "18da71384555cd",
		MessageTime: utils.GetNowStamp(),
		ImID:        IMID,
		Paid:        0,
		ImType:      8,
		Source:      "user",
		// imlpf4ff545dda0621f63e5d2f43cda1acfc
		SSID:        SSID,
		ContentType: "text",
		Content:     content,
		Trigger:     "input",
		UserID:      30301986,
		PVKey:       "b30c72f6dd4d8fd466963ea888" + "ff7d2f",
		WSID:        "b30c72f6dd4d8fd466963ea888" + "328382",
		Plat:        "PC",
		AppID:       "other",
	}
}

//
//// PhoneData Data 结构体表示 "data" 字段
//type PhoneData struct {
//	Seq         int64  `json:"seq"`
//	MessageID   string `json:"message_id"`
//	MsgID       string `json:"msgid"`
//	MessageTime string `json:"message_time"`
//	IMID        string `json:"imid"`
//	Paid        string `json:"paid"`
//	IMType      string `json:"im_type"`
//	Source      string `json:"source"`
//	SSID        string `json:"ssid"`
//	ContentType string `json:"content_type"`
//	Content     string `json:"content"`
//	Trigger     string `json:"trigger"`
//	FromID      string `json:"fromid"`
//	SourceTag   string `json:"source_tag"`
//	ShowType    string `json:"showType"`
//	UserID      string `json:"userid"`
//	PVKey       string `json:"pvKey"`
//	WSID        string `json:"wsid"`
//	Plat        string `json:"plat"`
//	AppID       string `json:"appid"`
//	Feature     string `json:"feature"`
//}
//
//// PhoneResp 结构体表示整个响应
//type PhoneResp struct {
//	Status  int       `json:"status"`
//	Data    PhoneData `json:"data"`
//	Message string    `json:"message"`
//	Feature struct {
//		PVKey         string   `json:"pvKey"`
//		ABTest        string   `json:"abTest"`
//		EffectiveTest bool     `json:"effectiveTest"`
//		Disables      []string `json:"disables"`
//		Enables       []string `json:"enables"`
//	} `json:"feature"`
//}
