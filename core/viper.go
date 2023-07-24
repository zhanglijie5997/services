package core

import (
	"fmt"
	"github.com/spf13/viper"
)

type MySqlConfig struct {
	Path         string `json:"path"`
	Port         string `json:"port"`
	Config       string `json:"config"`
	DbName       string `json:"db-name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	MaxIdleConns int64    `json:"max-idle-conns"`
	MaxOpenConns int64    `json:"max-open-conns"`
	LogMode      string `json:"log-mode"`
	LogZap       bool   `json:"log-zap"`
}

var SqlConfig MySqlConfig

func ViperInit()  {
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath("/etc/appname/")   // 查找配置文件所在的路径
	//viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
	viper.AddConfigPath(".")               // 还可以在工作目录中查找配置
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil { // 处理读取配置文件的错误
		fmt.Printf("Fatal error config file: %s \n", err)
	}
	_mysql := viper.Get("mysql").(map[string]interface{})

	//_mysql := mysql.(map[string]string)
	//_maxIdleConns, _ :=strconv.Atoi(_mysql["max-idle-conns"])
	//fmt.Println( _mysql)

	//
	SqlConfig = MySqlConfig{
		Path:         _mysql["path"].(string),
		Port:         _mysql["port"].(string),
		Config:       _mysql["config"].(string),
		DbName:       _mysql["db-name"].(string),
		Username:     _mysql["username"].(string),
		Password:     _mysql["password"].(string),
		MaxIdleConns: 10,
		MaxOpenConns: 0,
		LogMode:      "",
		LogZap:       false,
	}
	//// 设置配置文件的名字
	//viper.SetConfigName("config")
	//// 设置配置文件的类型
	//viper.SetConfigType("yaml")
	//// 添加配置文件的路径，指定 config 目录下寻找
	//viper.AddConfigPath("./")
	//// 寻找配置文件并读取
	//err := viper.ReadInConfig()
	//if err != nil {
	//	_ = fmt.Errorf("fatal error config file: %w", err)
	//}
	//mysql := viper.Get("mysql")
	//_mysql := mysql.(map[string]string)
	//_maxIdleConns, _err := strconv.ParseInt(string(_mysql["max-idle-conns"]), 10, 32)
	//if _err != nil {
	//	fmt.Printf("error")
	//	return ""
	//}
	//var config = MySqlConfig{
	//	Path:         _mysql["path"],
	//	Port:         _mysql["port"],
	//	Config:       _mysql["config"],
	//	DbName:       _mysql["db-name"],
	//	Username:     _mysql["username"],
	//	Password:     _mysql["password"],
	//	MaxIdleConns: _maxIdleConns,
	//	MaxOpenConns: 0,
	//	LogMode:      "",
	//	LogZap:       false,
	//}
	//SqlConfig = config
	//fmt.Println("db-name") // 127.0.0.1
	//return ""
}