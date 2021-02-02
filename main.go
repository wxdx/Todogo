package main

import (
	"Todogo/dao"
	"Todogo/routers"
	"Todogo/setting"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage：./Todogo conf/config.ini")
		return
	}
	// 加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
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
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
