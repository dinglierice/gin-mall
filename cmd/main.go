package main

import (
	"fmt"

	conf "mall/config"
	"mall/routes"

	_ "github.com/apache/skywalking-go"
)

func main() {
	Loading() // 加载配置
	r := routes.NewRouter()
	_ = r.Run(conf.Config.System.HttpPort)
	fmt.Println("启动配成功...")
}
