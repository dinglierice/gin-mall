package entcl

import (
	"context"
	"mall/ent"
	"mall/ent/psconfig"
	"mall/service"
	"strconv"
)

func GetPaginatedPsStrategy(ctx context.Context, limit, offset int) ([]*ent.PsStrategy, interface{}) {
	return entCli.PsStrategy.Query().
		Order(ent.Desc(psconfig.FieldCreateTime)).
		Limit(limit).
		Offset(offset).
		All(ctx)
}

func GetPsStrategy(ctx context.Context, id int) (*ent.PsStrategy, error) {
	return entCli.PsStrategy.Get(ctx, id)
}

func CreateStrategy(ctx context.Context, psStrategyService *service.PsStrategyService, ownerId uint) error {
	s := strconv.FormatUint(uint64(ownerId), 10) // 将uint转换为字符串
	i, err := strconv.ParseInt(s, 10, 64)        // 将字符串转换为int64
	if err != nil {
		return err
	}
	_, err = entCli.PsStrategy.Create().SetID(psStrategyService.ID).SetIsDelete(0).SetOwner(i).SetScriptContent(psStrategyService.ScriptContent).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
