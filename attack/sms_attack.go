package attack

import (
	"encoding/json"
	"goSMSBoom/log"
	"goSMSBoom/model"
	"goSMSBoom/utils"
)

// SMSAttack 处理短信发送请求
func SMSAttack(booms []*model.SMSBoom) error {
	for j := 0; j < len(booms); j++ {
		go func(j int) {
			boom := booms[j]
			desc := boom.Desc
			url := boom.Url
			method := boom.Method
			header := boom.Header
			data := boom.Data

			marshal, err := json.Marshal(data)
			if err != nil {
				log.Println("api.json中的", desc, "的data无法解析成为json")
				return
			}

			if method == "GET" {
				dos := utils.NewHttpDos(url, marshal, header)
				_, err = dos.Get()
				if err != nil {
					log.Println(desc, "get请求失败")
					return
				} else {
					log.Println("成功发送", desc, "短信")
				}
			} else if method == "POST" {
				dos := utils.NewHttpDos(url, marshal, header)
				_, err = dos.Post()
				if err != nil {
					log.Println(desc, "post请求失败")
					return
				} else {
					log.Println("成功发送", desc, "短信")
				}
			}
		}(j)
	}
	return nil
}
