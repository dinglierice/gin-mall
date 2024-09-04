package entcl

import (
	"context"
	"mall/ent"
	"mall/ent/psconfig"
	"mall/service"
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
	_, err := entCli.PsStrategy.Create().SetID(psStrategyService.ID).SetIsDelete(0).SetOwner(int(ownerId)).SetScriptContent(psStrategyService.ScriptContent).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
