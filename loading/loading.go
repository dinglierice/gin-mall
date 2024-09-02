package loading

import (
	util "mall/pkg/utils"
	"mall/repository/cache"
	"mall/repository/db/dao"
	"mall/repository/db/entcl"
)

func Loading() {
	// es.InitEs() // 如果需要接入ELK可以打开这个注释
	dao.InitMySQL()
	cache.InitCache()
	entcl.InitEntClient()
	// mq.InitRabbitMQ() // 如果需要接入RabbitMQ可以打开这个注释
	util.InitLog()
	go scriptStarting()
}

func scriptStarting() {
	// 启动一些脚本
}
