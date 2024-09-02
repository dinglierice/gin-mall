package entcl

import (
	"context"
	"log"
	"mall/ent"
)

var (
	entCli *ent.Client
)

func InitEntClient() {
	client, err := ent.Open("mysql", "root:123456@tcp(localhost:3306)/mall_db?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	// 运行数据库迁移
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	entCli = client
}
