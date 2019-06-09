package entity

import (
	"context"

	"github.com/yph152/api-frame-for-gin/internal/app/gindemo/schema"
	"github.com/yph152/api-frame-for-gin/pkg/gormplus"
)

// GetDemoDB 获取demo存储
func GetDemoDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	return getDBWithModel(ctx, defDB, Demo{})
}

// SchemaDemo demo
type SchemaDemo schema.Demo

// ToDemo 转换为demo实体
func (a SchemaDemo) ToDemo() *Demo {
	item := &Demo{
		RecordID: a.RecordID,
		Code:     a.Code,
		Name:     a.Name,
		Memo:     a.Memo,
		Status:   a.Status,
		Creator:  a.Creator,
	}
	return item
}

// Demo实体
type Demo struct {
	Model
	RecordID string `gorm:"column:record_id;"`
	Code     string `gorm:"column:code;"`
	Name     string `gorm:"column:name;"`
	Memo     string `gorm:"column:memo;"`
	Status   int    `gorm:"column:status;'`
	Creator  string `gorm:"column:creator;"`
}

func (a Demo) String() string {
	return toString(a)
}

// TableName
func (a Demo) TableName() string {
	return a.Model.TableName("demo")
}

// ToSchemaDemo 转换为demo对象
func (a Demo) ToSchemaDemo() *schema.Demo {
	item := &schema.Demo{
		RecordID:  a.RecordID,
		Code:      a.Code,
		Name:      a.Name,
		Memo:      a.Memo,
		Status:    a.Status,
		Creator:   a.Creator,
		CreatedAt: a.CreatedAt,
	}
	return item
}

// Demos demo列表
type Demos []*Demo

// ToSchemaDemos 转换为demo对象列表
func (a Demos) ToSchemaDemos() []*schema.Demo {
	list := make([]*schema.Demo, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaDemo()
	}
	return list
}
