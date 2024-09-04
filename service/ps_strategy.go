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

type PsStrategyService struct {
	ID            int       `json:"id"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
	Owner         int       `json:"owner"`
	ScriptContent string    `json:"script_content"`
	IsDelete      int       `json:"is_delete"`
	model2.BasePage
}

func (service *PsStrategyService) List(ctx context.Context) any {
	var configs []*ent.PsStrategy

	if service.PageSize == 0 {
		service.PageSize = 15
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		configs, _ = entcl.GetPaginatedPsStrategy(ctx, service.PageSize, service.PageNum)
		wg.Done()
	}()
	wg.Wait()
	return serializer.BuildListResponse(serializer.BuildPsStrategy4Ent(configs), uint(len(configs)))
}

func (service *PsStrategyService) Show(ctx context.Context, param string) interface{} {
	id, _ := strconv.Atoi(param)
	code := e.SUCCESS
	psStrategy, err := entcl.GetPsStrategy(ctx, id)
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
		Data:   serializer.BuildPsStrategy(psStrategy),
		Msg:    e.GetMsg(code),
	}
}

func (service *PsStrategyService) Create(ctx context.Context, createPsStrategyService *PsStrategyService, claimId uint) serializer.Response {
	var _ *ent.PsStrategy
	var err error
	code := e.SUCCESS

	// 判断有无这个商品
	_, err = entcl.GetPsStrategy(ctx, service.ID)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 创建购物车
	err = entcl.CreateStrategy(ctx, createPsStrategyService, claimId)
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
		Msg:    e.GetMsg(code),
	}
}
