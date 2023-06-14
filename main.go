package main

import "gitee.com/lybbn/go-vue-lyadmin/utils/cmd"

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server
// @securityDefinitions.apikey  JWT (apiKey)
// @in                          header
// @name                        Authorization
// @BasePath                    /
func main() {
	cmd.Execute()
}
