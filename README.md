# mygin

学习gin

```shell
# 初始化go.mod
go mod init mygin  
 
# 下载并安装 gin
go get -u github.com/gin-gonic/gin
```


```shell
# 第三方工具 air 的 github 地址：https://github.com/air-verse/air
# 下载 第三方热加载工具 air(用于热加载启动项目)
go install github.com/air-verse/air@latest

# You can initialize the .air.toml configuration file to the current directory with the default settings running the following command.
air init

# After this, you can just run the air command without additional arguments, and it will use the .air.toml file for configuration.
air

#热启动服务并设置端口，命令行执行: 
# mac下
PORT=8086 air
# windows下 
$env:PORT="8086"
air
```