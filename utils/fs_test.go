package utils

import (
	"fmt"
	l "goSMSBoom/log"
	"log"
	"os"
	"testing"
)

func TestGetFileByteArr(t *testing.T) {
	arr, err := GetFileByteArr("api.json")
	if err != nil {
		l.Println(err)
	}
	l.Println(string(arr))

}

func TestH(t *testing.T) {
	file := "./" + "message" + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[qSkipTool]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	log.Println("Hello Davis!") // log 还是可以作为输出的前缀
}

func TestOne(t *testing.T) {
	//https://jdapi.jd100.com/uc/v1/getSMSCode?account=18248625125&sign_type=1&use_type=1

	dos := NewHttpDos("https://jdapi.jd100.com/uc/v1/getSMSCode?account=17757488665&sign_type=1&use_type=1",
		nil,
		nil)
	get, err := dos.Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(get))

}
