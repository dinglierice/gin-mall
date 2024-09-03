package serializer

import (
	"mall/ent"
	"time"
)

type PsStrategy struct {
	ID            int       `json:"id"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
	Owner         int       `json:"owner"`
	ScriptContent string    `json:"script_content"`
	IsDelete      int       `json:"is_delete"`
}

func BuildPsStrategy(item *ent.PsStrategy) PsStrategy {
	return PsStrategy{
		ID:            item.ID,
		CreateTime:    item.CreateTime,
		UpdateTime:    item.UpdateTime,
		Owner:         item.Owner,
		ScriptContent: item.ScriptContent,
		IsDelete:      item.IsDelete,
	}
}

func BuildPsStrategy4Ent(items []*ent.PsStrategy) (configs []PsStrategy) {
	for _, item := range items {
		config := BuildPsStrategy(item)
		configs = append(configs, config)
	}
	return configs
}
