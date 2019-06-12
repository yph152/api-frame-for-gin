package gorm

import (
	"github.com/yph152/api-frame-for-gin/internal/app/gindemo/model"
	"github.com/yph152/api-frame-for-gin/pkg/gormplus"

	"github.com/yph152/api-frame-for-gin/internal/app/gindemo/model/gorm/entity"
	gmodel "github.com/yph152/api-frame-for-gin/internal/app/gindemo/model/gorm/model"
)

// SetTablePrefix 设定表名前缀
func SetTablePrefix(prefix string) {
	entity.SetTablePrefix(prefix)
}

// AutoMigrate 自动映射数据表
func AutoMigrate(db *gormplus.DB) error {
	return db.AutoMigrate(
		new(entity.Demo),
	).Error
}

// NewModel 创建gorm存储，实现统一的存储接口
func NewModel(db *gormplus.DB) *model.Common {
	return &model.Common{
		Trans: gmodel.NewTrans(db),
		Demo:  gmodel.NewDemo(db),
	}
}
