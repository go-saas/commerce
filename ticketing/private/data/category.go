package data

import (
	"context"
	"errors"
	"fmt"
	v12 "github.com/go-saas/commerce/ticketing/api/category/v1"
	"github.com/go-saas/commerce/ticketing/private/biz"
	kitgorm "github.com/go-saas/kit/pkg/gorm"
	"github.com/go-saas/kit/pkg/query"
	sgorm "github.com/go-saas/saas/gorm"
	"github.com/goxiaoy/go-eventbus"
	"gorm.io/gorm"
	"strings"
)

type CategoryRepo struct {
	*kitgorm.Repo[biz.TicketingCategory, string, v12.ListCategoryRequest]
}

func NewCategoryRepo(dbProvider sgorm.DbProvider, eventbus *eventbus.EventBus) biz.TicketingCategoryRepo {
	res := &CategoryRepo{}
	res.Repo = kitgorm.NewRepo[biz.TicketingCategory, string, v12.ListCategoryRequest](dbProvider, eventbus, res)
	return res
}

func (c *CategoryRepo) GetDb(ctx context.Context) *gorm.DB {
	return GetDb(ctx, c.Repo.DbProvider)
}

func (c *CategoryRepo) BuildPrimaryField() string {
	return "key"
}

func preloadParent(maxDepth, currentDepth int) func(d *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if currentDepth > maxDepth {
			return d
		}
		return d.Preload("Parent", preloadParent(maxDepth, currentDepth+1))
	}
}

// BuildFilterScope filter
func (c *CategoryRepo) BuildFilterScope(q *v12.ListCategoryRequest) func(db *gorm.DB) *gorm.DB {
	search := q.Search
	filter := q.Filter
	return func(db *gorm.DB) *gorm.DB {
		ret := db

		if search != "" {
			ret = ret.Where("name like ?", fmt.Sprintf("%%%v%%", search))
		}
		if filter == nil {
			return ret
		}

		if filter.Name != nil {
			ret = ret.Scopes(kitgorm.BuildStringFilter("`name`", filter.Name))
		}
		if filter.Parent != nil {
			ret = ret.Scopes(kitgorm.BuildStringFilter("`parent_id`", filter.Name))
		}
		return ret
	}
}

func (c *CategoryRepo) DefaultSorting() []string {
	return []string{"created_at"}
}

func (c *CategoryRepo) Create(ctx context.Context, entity *biz.TicketingCategory) error {
	//set path
	err := c.setPath(ctx, entity)
	if err != nil {
		return err
	}
	return c.Repo.Create(ctx, entity)
}

func (c *CategoryRepo) Update(ctx context.Context, id string, entity *biz.TicketingCategory, p query.Select) error {
	oldPath := entity.Path
	allChildren, err := c.FindAllChildren(ctx, entity)
	if err != nil {
		return err
	}
	//set path
	err = c.setPath(ctx, entity)
	if err != nil {
		return err
	}
	err = c.Repo.Update(ctx, id, entity, p)
	if err != nil {
		return err
	}
	for _, child := range allChildren {
		//set new path
		child.Path = entity.Path + "/" + strings.TrimPrefix(child.Path, oldPath)
		err = c.Repo.Update(ctx, id, child, p)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CategoryRepo) FindAllChildren(ctx context.Context, entity *biz.TicketingCategory) ([]*biz.TicketingCategory, error) {
	var children []*biz.TicketingCategory
	err := c.GetDb(ctx).Model(&biz.TicketingCategory{}).Where("path LIKE ?", entity.Path+"/").Find(&children).Error
	if err != nil {
		return nil, err
	}
	return children, nil
}

func (c *CategoryRepo) setPath(ctx context.Context, entity *biz.TicketingCategory) error {
	if entity.ParentID != nil {
		var parent = &biz.TicketingCategory{}
		err := c.GetDb(ctx).Model(&biz.TicketingCategory{}).Scopes(preloadParent(100, 0)).Where("id = ?", *entity.ParentID).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			return err
		}
		var key []string
		for parent != nil {
			key = append([]string{parent.Key}, key...)
			parent = parent.Parent
		}
		entity.Path = strings.Join(key, "/")
	}
	return nil
}
