package test

import (
	"encoding/json"
	"goSMSBoom/model"
	"goSMSBoom/utils"
	"testing"
)

func TestHttpDos_Get(t *testing.T) {
	get, err := utils.NewHttpDos("https://ada.baidu.com/imlp-message/imlp/send?seq=1&message_id=key18da6d12a6fe319&msgid=key18da6d12a6fe319&message_time=1707900742255&imid=9a77260b6026231a457349a4f46bf074&paid=0&im_type=8&source=user&ssid=imlp3b400e4ed2497b6fb01159f5e44d2aeb&content_type=text&content=%E4%BD%A0%E5%A5%BD&trigger=input&fromid=&source_tag=&showType=&userid=41854658&pvKey=dc28eefbce06bccea9245d281bceff4e&wsid=01b68addc8150b48551e674eda741989&plat=PC&appid=other&feature=",
		nil, nil).Get()
	if err != nil {
		t.Error(err)
	}
	var getRes model.PhoneResp
	err = json.Unmarshal(get, &getRes)
	if err != nil {
		t.Error(err)
	}
	t.Log("=====", getRes)
	if getRes.Status != 0 {
		t.Error("get请求失败")
	}

}
