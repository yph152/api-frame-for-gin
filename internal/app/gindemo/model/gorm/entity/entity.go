package entity

import (
	"context"
	"fmt"

	"github.com/yph152/api-frame-for-gin/pkg/util"

	icontext "github.com/yph152/api-frame-for-gin/internal/app/gindemo/context"
	"github.com/yph152/api-frame-for-gin/pkg/gormplus"
)

// 表名前缀
var tablePrefix string

// SetTablePrefix 设置表名前缀
func SetTablePrefix(prefix string) {
	tablePrefix = prefix
}

// GetTablePrefix 获取表名前缀
func GetTablePrefix() string {
	return tablePrefix
}

type Model struct {
	ID        string `gorm:"column:id;primary_key;"`
	CreatedAt int64  `gorm:"column:created_at;"`
	UpdatedAt int64  `gorm:"column:updated_at;"`
	DeletedAt int64  `gorm:"column:deleted_at;"`
}

// TableName table name
func (Model) TableName(name string) string {
	return fmt.Sprintf("%s%s", GetTablePrefix(), name)
}

func toString(v interface{}) string {
	return util.JSONMarshalToString(v)
}

func getDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	trans, ok := icontext.FromTrans(ctx)
	if ok {
		db, ok := trans.(*gormplus.DB)
		if ok {
			return db
		}
	}
	return defDB
}

func getDBWithModel(ctx context.Context, defDB *gormplus.DB, m interface{}) *gormplus.DB {
	return gormplus.Wrap(getDB(ctx, defDB).Model(m))
}
