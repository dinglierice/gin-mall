package entcl

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"log"
	"mall/ent"
)

var (
	entCli *ent.Client
)

func InitEntClient() {
	drv, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/mall_db?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	// 使用自定义日志驱动包装原始驱动
	logger := EntLogger{}
	logDrv := dialect.DebugWithContext(drv, logger.Log)

	// 使用包装后的驱动创建客户端
	client := ent.NewClient(ent.Driver(logDrv))

	// 运行数据库迁移
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	entCli = client
}
