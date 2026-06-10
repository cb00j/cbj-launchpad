package models

import "time"

type SyncState struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"column:name;uniqueIndex"` // 标识,如 "register_listener"
	LastBlock uint64 `gorm:"column:last_block"`       // 最后处理的区块
	UpdatedAt time.Time
}

func (SyncState) TableName() string {
	return "sync_state"
}
