#### go build

`go build -n main.go` 打印编译过程且不执行编译后的二进制文件 

`go build -x main.go` 打印编译过程且执行编译后的二进制文件

`go build` 在当前目录生成一个与目录名一致的二进制执行文件

`go build -o blog-service` 指定编译后的二进制执行文件名

`go install` 编译后安装到$GOBIN目录

`CGO_ENABLE=0 GOOS=linux go build -a -o blog-service` 交叉编译生成指定平台[linux]的可执行文件

`go build -ldflags="-w -s"`缩小生成的二进制文件 -w 去除DWARF调试信息：panic无文件名、行号 -s 去除符号表：无法gdb调试

`upx blog-service` 对二进制文件压缩

`go build -ldflags "-X main.appName=Learn Golang"` 设置编译信息

go build -ldflags "-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.0.0 -X main.gitCommitID=`git rev-parse HEAD`"

####  重启和停止

`kill -l` 系统支持的所有信号
