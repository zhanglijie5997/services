package core

import "fmt"

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap() string {
	fmt.Println("test Zap")
	return ""
	//if ok, _ := utils.PathExists(global.GVA_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
	//	fmt.Printf("create %v directory\n", global.GVA_CONFIG.Zap.Director)
	//	_ = os.Mkdir(global.GVA_CONFIG.Zap.Director, os.ModePerm)
	//}
	//
	//cores := internal.Zap.GetZapCores()
	//logger = zap.New(zapcore.NewTee(cores...))
	//
	//if global.GVA_CONFIG.Zap.ShowLine {
	//	logger = logger.WithOptions(zap.AddCaller())
	//}
	//return logger
}
