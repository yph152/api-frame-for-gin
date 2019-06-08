package model

import (
	"context"
)

type Trans interface {
	// 开始事务
	Begin(ctx context.Context) (interface{}, error)

	// 提交事务
	Commit(ctx context.Context, trans interface{}) error

	// 回滚事务
	Rollback(ctx context.Context, trans interface{}) error
}
