#!/bin/bash
echo "go build"
go generate
go build -o golyadmin main.go
chmod +x ./golyadmin
echo "kill golyadmin service"
killall golyadmin # kill golyadmin service
nohup ./golyadmin start >> access.log 2>&1 & #后台启动服务将日志写入access.log文件
echo "run golyadmin success"
ps -aux | grep golyadmin
