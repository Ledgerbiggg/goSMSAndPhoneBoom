package log

import (
	"fmt"
	"goSMSBoom/config"
	"log"
	"os"
	"strings"
	"time"
)

// InitLogStyle 初始化日志格式
func InitLogStyle() {
	var fileName string
	if config.Configs.ENV == "win" {
		// 获取当前时间并格式化为文件名
		format := time.Now().Format("2006-01-02 15:04:05")
		// 将 ":" 替换为其他合法字符，比如 "-"
		format = strings.Replace(format, ":", "-", -1)
		fileName = format + ".txt"
	} else {
		fileName = "log.txt"
	}
	fmt.Println("日志地址:" + fileName)

	// 打开日志文件，如果文件不存在则创建
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}

	// 设置日志的输出目标为文件
	log.SetOutput(file)
	log.SetPrefix("[qSkipTool]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}

// Println 打印
func Println(a ...any) {
	// 控制台输出日志
	log.Println(a...)
	fmt.Println(a)
}
