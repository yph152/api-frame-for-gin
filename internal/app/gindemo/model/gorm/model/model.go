package model

import (
	"context"

	icontext "github.com/yph152/api-frame-for-gin/internal/app/gindemo/context"
	"github.com/yph152/api-frame-for-gin/internal/app/gindemo/schema"
	"github.com/yph152/api-frame-for-gin/pkg/gormplus"
)

// ExecTrans 执行事务
func ExecTrans(ctx context.Context, db *gormplus.DB, fn func(context.Context) error) error {
	if _, ok := icontext.FromTrans(ctx); ok {
		return fn(ctx)
	}

	transModel := NewTrans(db)
	trans, err := transModel.Begin(ctx)
	if err != nil {
		return err
	}

	err = fn(icontext.NewTrans(ctx, trans))
	if err != nil {
		_ = transModel.Rollback(ctx, trans)
		return err
	}
	return transModel.Commit(ctx, trans)
}

// WrapPageQuery 包装带有分页的查询
func WrapPageQuery(db *gorm.DB, pp *schema.PaginationParam, out interface{}) (*schema.Paginationresult, error) {
	if pp != nil {
		total, err := gormplus.Wrap(db).FindPage(db, pp.PageIndex, pp.PageSize, out)
		if err != nil {
			return nil, err
		}
		return &schema.PaginationResult{
			Total: total,
		}, nil
	}

	result := db.Find(out)
	return nil, result.Error
}
