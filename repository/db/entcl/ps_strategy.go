package entcl

import (
	"context"
	"mall/ent"
	"mall/ent/psconfig"
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
