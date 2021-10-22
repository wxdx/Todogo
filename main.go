package main

import (
	"Todogo/dao"
	"Todogo/routers"
	"Todogo/setting"
	"fmt"
	"log"
)

func main() {

	//if len(os.Args) < 2 {
	//	fmt.Println("Usage：./Todogo conf/config.ini")
	//	return
	//}
	// 加载配置文件 os.Args[1]
	if err := setting.Init("conf/config.ini"); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}

	//初始化数据库
	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		panic(err)
	}
	defer dao.Db.Close()

	//注册路由
	app := routers.SetupRouter()
	log.Fatal(app.Listen(fmt.Sprintf(":%d", setting.Conf.Port)))
}
