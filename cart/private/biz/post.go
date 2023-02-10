package biz

import (
	"github.com/go-saas/kit/pkg/data"
	"github.com/go-saas/kit/pkg/gorm"
)

type Post struct {
	gorm.UIDBase
	gorm.AuditedModel
	Name string
}

type PostRepo interface {
	data.Repo[Post, string, interface{}]
}
