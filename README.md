## 使用
- 找到cmd下的main函数初始化项目
```go
go build main.go
```

## 包说明
- api 
    - 放置swagger相关的一些docs文件
    
- app
    - 项目主要业务所在包
    
- assets
    - 静态资源包
    
- cmd
    - 项目启动main函数所在目录
    
- git hook
    - git hook脚本
    
- library
    - 业务公共组件包,以后可直接打包带走

- script
    - 可能的脚本放置地
    
- tool
    - 工具包