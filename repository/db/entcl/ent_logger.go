package entcl

import (
	"context"
	"fmt"
	"log"
	"time"
)

// EntLogger 是一个自定义的Ent日志驱动
type EntLogger struct{}

func (EntLogger) Log(ctx context.Context, v ...any) {
	prefix := fmt.Sprintf("ENT_SQL: %s", time.Now().Format("2006-01-02 15:04:05"))
	if len(v) > 0 {
		if s, ok := v[0].(string); ok {
			sql := fmt.Sprintf(s, v[1:]...)
			log.Printf("%s: %s\n", prefix, sql)
		}
	}
}
