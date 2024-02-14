/*
@author: ledger
@since: 2024/1/29
*/

package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

// Configs 使用全局的配置变量
var Configs *PhoneBoom

// PhoneBoom 结构体用于表示手机号轰炸配置
type PhoneBoom struct {
	ENV           string `yaml:"ENV"`
	Phone         string `yaml:"Phone"`
	ThreadCount   int    `yaml:"ThreadCount"`
	ExecutionCron string `yaml:"ExecutionCron"`
	Content       string `yaml:"Content"`
	SSID          string `yaml:"SSID"`
}

// LoadConfig viper读取yaml
func LoadConfig() error {
	// yaml
	vconfig := viper.New()
	//表示 先预加载匹配的环境变量
	vconfig.AutomaticEnv()
	//设置环境变量分割符，将点号和横杠替换为下划线
	vconfig.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	// 设置读取的配置文件
	vconfig.SetConfigName("config")
	// 添加读取的配置文件路径
	vconfig.AddConfigPath(".")
	// 读取文件类型
	vconfig.SetConfigType("yaml")
	err := vconfig.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	if err = vconfig.Unmarshal(&Configs); err != nil {
		log.Panicln("unmarshal cng file fail " + err.Error())
	}
	// 赋值全局变量
	return err
}
