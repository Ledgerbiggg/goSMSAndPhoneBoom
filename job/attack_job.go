package job

import (
	"goSMSBoom/attack"
	"goSMSBoom/config"
	"goSMSBoom/log"
)

func AttackJob() error {

	// 创建一个新的 cron 实例
	c := NewWithSeconds()

	// 准备
	err := attack.Start()
	if err != nil {
		log.Println("配置文件加载失败")
		return err
	}

	_, err = c.AddFunc(config.Configs.ExecutionCron, func() {
		for i := 0; i < config.Configs.ThreadCount; i++ {
			err = attack.StartBoom()
			if err != nil {
				panic("启动失败")
			}
		}
	})
	if err != nil {
		log.Println("定时器启动失败", err)
		return err
	}

	// 启动 cron 服务
	c.Start()

	// 阻塞主线程
	select {}
}
