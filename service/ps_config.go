package service

import (
	"context"
	logging "github.com/sirupsen/logrus"
	"mall/ent"
	"mall/pkg/e"
	"mall/repository/db/entcl"
	model2 "mall/repository/db/model"
	"mall/serializer"
	"strconv"
	"sync"
	"time"
)

type PsConfigService struct {
	PsID       int       `json:"ps_id"`       // 主键
	PsScene    string    `json:"ps_scene"`    // 场景名称
	PsFilter   int       `json:"ps_filter"`   // 过滤策略
	PsMessage  int       `json:"ps_message"`  // 消息策略，可能为null
	PsEvent    int       `json:"ps_event"`    // 时间策略，可能为null
	PsFeature  int       `json:"ps_feature"`  // 特征策略，可能为null
	PsStrategy int       `json:"ps_strategy"` // 脚本，可能为null
	OwnerID    int       `json:"owner_id"`    // 所有者ID，可能为null
	Managers   string    `json:"managers"`    // 管理员列表，可能为null
	UpdateUser int       `json:"update_user"` // 更新用户，可能为null
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
	model2.BasePage
}

func (service PsConfigService) List(ctx context.Context) interface{} {
	var configs []*ent.PsConfig

	if service.PageSize == 0 {
		service.PageSize = 15
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		configs, _ = entcl.GetPaginatedPsConfigs(ctx, service.PageSize, service.PageNum)
		wg.Done()
	}()
	wg.Wait()
	return serializer.BuildListResponse(serializer.BuildPsConfigs4Ent(configs), uint(len(configs)))
}

func (service PsConfigService) Show(ctx context.Context, param string) interface{} {
	id, _ := strconv.Atoi(param)
	code := e.SUCCESS
	config, err := entcl.GetPsConfig(ctx, id)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildPsConfig(config),
		Msg:    e.GetMsg(code),
	}
}
