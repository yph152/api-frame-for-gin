package gindemo

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/yph152/api-frame-for-gin/internal/app/gindemo/config"
	"github.com/yph152/api-frame-for-gin/internal/app/gindemo/model"
	"github.com/yph152/api-frame-for-gin/internal/app/gindemo/model/gorm"
	"github.com/yph152/api-frame-for-gin/pkg/gormplus"
)

// InitStore 初始化存储，返回统一的存储接口
func InitStore() (*model.Common, func(), error) {
	var (
		storeCall func()
		m         *model.Common
	)

	cfg := config.GetGlobalConfig()
	fmt.Println(cfg)
	switch cfg.Store {
	case "gorm":
		db, err := initGorm()
		if err != nil {
			return nil, nil, err
		}

		storeCall = func() {
			db.Close()
		}

		gorm.SetTablePrefix(cfg.Gorm.TablePrefix)
		err = gorm.AutoMigrate(db)
		if err != nil {
			return nil, nil, err
		}
		m = gorm.NewModel(db)
	default:
		return nil, nil, errors.New("unknow store")
	}

	return m, storeCall, nil
}

// initGorm 实例化gorm存储
func initGorm() (*gormplus.DB, error) {
	cfg := config.GetGlobalConfig()

	var dsn string
	switch cfg.Gorm.DBType {
	case "mysql":
		dsn = cfg.MySQL.DSN()
	case "sqlite3":
		dsn = cfg.Sqlite3.DSN()
		os.MkdirAll(filepath.Dir(dsn), 0777)
	case "postgres":
		dsn = cfg.Postgres.DSN()
	default:
		return nil, errors.New("unknow db")
	}

	return gormplus.New(&gormplus.Config{
		Debug:        cfg.Gorm.Debug,
		DBType:       cfg.Gorm.DBType,
		DSN:          dsn,
		MaxIdleConns: cfg.Gorm.MaxIdleConns,
		MaxLifetime:  cfg.Gorm.MaxLifetime,
		MaxOpenConns: cfg.Gorm.MaxOpenConns,
	})
}
