package main

import (
	"fmt"
	"service/core"
	dbInit "service/db"
	"service/router"
)

func main() {
	//global.LOG = zip.Zap()
	core.Zap()
	core.ViperInit()
	dbInit.Db()
	fmt.Println("logs", core.SqlConfig)
	router.Router()
}
