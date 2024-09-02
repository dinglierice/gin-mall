package entcl

import (
	"context"
	"mall/ent"
	"mall/ent/psconfig"
)

func GetPaginatedPsConfigs(ctx context.Context, limit, offset int) ([]*ent.PsConfig, error) {
	return entCli.PsConfig.Query().
		Order(ent.Desc(psconfig.FieldCreateTime)).
		Limit(limit).
		Offset(offset).
		All(ctx)
}

func GetPsConfig(ctx context.Context, id int) (*ent.PsConfig, error) {
	return entCli.PsConfig.Get(ctx, id)
}
