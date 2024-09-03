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

func (service *PsStrategyService) Create(ctx context.Context, id uint) serializer.Response {
	//var user *ent.PsStrategy
	//var err error
	//code := e.SUCCESS
	// 找到用户
	//userDao := dao2.NewUserDao(ctx)
	//user, err = userDao.GetUserById(uId)
	//if err != nil {
	//	logging.Info(err)
	//	code = e.ErrorDatabase
	//	return serializer.Response{
	//		Status: code,
	//		Msg:    e.GetMsg(code),
	//		Error:  err.Error(),
	//	}
	//}
	//if service.NickName != "" {
	//	user.NickName = service.NickName
	//}
	//
	//err = userDao.UpdateUserById(uId, user)
	//if err != nil {
	//	logging.Info(err)
	//	code = e.ErrorDatabase
	//	return serializer.Response{
	//		Status: code,
	//		Msg:    e.GetMsg(code),
	//		Error:  err.Error(),
	//	}
	//}
	//
	//return serializer.Response{
	//	Status: code,
	//	Data:   serializer.BuildUser(user),
	//	Msg:    e.GetMsg(code),
	//}
	return serializer.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
	}
}
