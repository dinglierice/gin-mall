package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"mall/conf"
	_ "mall/docs" // 这里导入生成的docs包
	"mall/loading"
	"mall/routes"
)

func main() {
	conf.Init()
	loading.Loading()
	r := routes.NewRouter()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	_ = r.Run(conf.HttpPort)
}
