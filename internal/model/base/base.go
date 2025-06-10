package base

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time `gorm:"autoCreateTime"` // 指定为自动创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // 指定为自动更新时间
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	// 为 ID 字段生成一个新的 UUID
	base.Id = uuid.New()
	return
}
