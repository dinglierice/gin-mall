package serializer

import (
	"mall/ent"
	"mall/repository/db/model"
	"time"
)

type PsConfig struct {
	PsID       int       `json:"ps_id"`
	PSName     string    `json:"ps_scene"`
	PSFilter   int       `json:"ps_filter"`
	PSMessage  *int      `json:"ps_message"`
	PSEvent    *int      `json:"ps_event"`
	PSFeature  *int      `json:"ps_feature"`
	PSStrategy *int      `json:"ps_strategy"`
	OwnerID    *int      `json:"owner_id"`
	Managers   *string   `json:"managers"`
	UpdateUser *int      `json:"update_user"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func BuildPsConfig(item *model.PsConfig) PsConfig {
	return PsConfig{
		PsID:       item.PsID,
		PSName:     item.PsScene,
		PSFilter:   item.PsFilter,
		PSMessage:  item.PsMessage,
		PSEvent:    item.PsEvent,
		PSFeature:  item.PsFeature,
		PSStrategy: item.PsStrategy,
		OwnerID:    item.OwnerID,
		Managers:   item.Managers,
	}
}

func BuildPsConfig2(item *ent.PsConfig) PsConfig {
	return PsConfig{
		PsID:       item.ID,
		PSName:     item.PsScene,
		PSFilter:   item.PsFilter,
		PSMessage:  item.PsMessage,
		PSEvent:    item.PsEvent,
		PSFeature:  item.PsFeature,
		PSStrategy: item.PsStrategy,
		OwnerID:    item.OwnerID,
		Managers:   item.Managers,
	}
}

func BuildPsConfigs(items []*model.PsConfig) (configs []PsConfig) {
	for _, item := range items {
		config := BuildPsConfig(item)
		configs = append(configs, config)
	}
	return configs
}

func BuildPsConfigs4Ent(items []*ent.PsConfig) (configs []PsConfig) {
	for _, item := range items {
		config := BuildPsConfig2(item)
		configs = append(configs, config)
	}
	return configs
}