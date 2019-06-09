/*
Package ginadmin 生成swagger文档
文档规则请参考：https://github.com/teambition/swaggo/wiki/Declarative-Comments-Format
使用方式：
	go get -u -v github.com/teambition/swaggo
		swaggo -s ./internal/app/ginadmin/swagger.go -p . -o ./internal/app/ginadmin/swagger
*/
package gindemo

import (
	// API控制器
	_ "github.com/yph152/api-frame-for-gin/internal/app/gindemo/routers/apis/ctl"
)

// @Version 3.1.1
// @Title GinDemo
// @Description RBAC scaffolding based on GIN + GORM + CASBIN.
// @Schemes http,https
// @Host 127.0.0.1:10088
// @BasePath /
// @Name yph152
// @Contact yph152@gmail.com
// @Consumes json
// @Produces json
