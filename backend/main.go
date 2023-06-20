package main

import "gitee.com/lybbn/go-vue-lyadmin/utils/cmd"

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger API
// @version                     0.0.1
// @description                 This is a go-vue-lyadmin Server
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @BasePath                    /api
func main() {
	cmd.Execute()
}
