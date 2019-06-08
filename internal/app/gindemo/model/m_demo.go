package model

import (
	"context"

	"github.com/yph152/api-frame-for-gin/internal/app/gindemo/schema"
)

type Demo interface {
	// 查询接口
	Query(ctx context.Context, params schema.DemoQueryParam, opts ...schema.DemoQueryOptions) (*schema.DemoQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string, opts ...schema.DemoQueryOptions) (*schema.Demo, error)
	// 创建数据
	Create(ctx context.Context, item schema.Demo) error
	// 更新数据
	Update(ctx context.Context, recordID string, item schema.Demo) error
	// 删除数据
	Delete(ctx context.Context, recordID string) error
}
