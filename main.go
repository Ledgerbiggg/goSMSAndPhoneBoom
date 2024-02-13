package main

import (
	"encoding/json"
	"goSMSBoom/config"
	"goSMSBoom/job"
	"goSMSBoom/log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	err := config.LoadConfig()
	if err != nil {
		log.Println("配置文件加载失败")
	}
	marshal, err := json.Marshal(config.Configs)
	log.Println("配置文件是", string(marshal))
	log.InitLogStyle()
}

func main() {
	log.Println("开始攻击")
	err := job.AttackJob()
	log.Println("定时器启动失败", err)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case <-interrupt:
		log.Println("结束攻击")
	}
}
