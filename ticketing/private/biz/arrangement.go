package biz

import (
	kitgorm "github.com/go-saas/kit/pkg/gorm"
	"time"
)

type Arrangement struct {
	kitgorm.UIDBase

	ActivityID string
	Activity   Activity `gorm:"foreignKey:ActivityID"`

	StartTime time.Time

	LocationID string
	Location   Location `gorm:"foreignKey:LocationID"`

	Halls []Hall `gorm:"many2many:arrangement_halls;"`
}
