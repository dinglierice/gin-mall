package model

import "time"

type PsConfig struct {
	PsID       int        `gorm:"primaryKey;autoIncrement;comment:'主键'" json:"ps_id"`
	PsScene    string     `gorm:"type:varchar(255);not null;uniqueIndex:ps_scene;comment:'场景名称'" json:"ps_scene"`
	PsFilter   int        `gorm:"not null;comment:'过滤策略'" json:"ps_filter"`
	PsMessage  *int       `gorm:"comment:'消息策略'" json:"ps_message"`
	PsEvent    *int       `gorm:"comment:'时间策略'" json:"ps_event"`
	PsFeature  *int       `gorm:"comment:'特征策略'" json:"ps_feature"`
	PsStrategy *int       `gorm:"comment:'脚本'" json:"ps_strategy"`
	OwnerID    *int       `gorm:"comment:'所有者ID'" json:"owner_id"`
	Managers   *string    `gorm:"type:varchar(255);comment:'管理员列表'" json:"managers"`
	UpdateUser *int       `gorm:"comment:'更新用户'" json:"update_user"`
	CreateTime *time.Time `gorm:"comment:'创建时间'" json:"create_time"`
	UpdateTime *time.Time `gorm:"comment:'更新时间'" json:"update_time"`
}
