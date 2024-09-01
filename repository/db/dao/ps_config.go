package dao

import (
	"context"
	model2 "mall/repository/db/model"

	"gorm.io/gorm"
)

type PsConfigDao struct {
	*gorm.DB
}

func NewPsConfigDao(ctx context.Context) *PsConfigDao {
	return &PsConfigDao{NewDBClient(ctx)}
}

func (dao *PsConfigDao) ListPsConfigByCondition(condition map[string]interface{}, page model2.BasePage) (products []*model2.PsConfig, err error) {
	err = dao.DB.Where(condition).
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&products).Error
	return
}
