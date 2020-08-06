
## 格式规范

#### 文件夹名称 和 package包名

* 只包含小写字母a-z，如 api, cmd, githook

#### .go文件包引用格式

* 按标准包、第三方包、项目内包分三部分，每一部分用空行分隔，标准包和第三方包尽量避免使用省略包名(dot import alias)，必要时使用小写字母定义包别名

```go
package endpoint

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"micro_service/internal/app/client/model"
)
```

* 只包含小写字母a-z，如 api, cmd, githook

#### 代码文件名称

* 只包含小写字母a-z和下划线_，如 [rpc_client.go](https://gitee.com/Skyd188/micro_services/blob/master/internal/app/client/myendpoint/rpc_client.go)

#### 函数/方法名称

* 大小写字母a-z A-Z，通常避免使用特殊符号和数字，考虑使用`动词+名词`规则

```go
package endpoint

func MakeEchoEndpoint(svc pb.MyServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.EchoRequest)
		return svc.Echo(ctx, req)
	}
}
```

#### 常量定义

* 驼峰命名，且通常在文件开头统一声明

```go
package user

import "github.com/go-kit/kit"

const (
    userName    = "Tom" //unexported
    CompanyName = "Meross" //exported
)
```

#### 变量定义

* 驼峰命名，且函数内变量应尽量在函数开头统一声明

```go
package demo

func DoSomeThing(str string) error {
    var user User

    if user, ok := userMap[str]; !ok {
        return errors.New("not found")
    }
    
    fmt.Println(user)
    return nil
}
```

#### 函数返回值

* 编写工具方法时，可以事先声明函数返回字段名称，IDE会在`GetContext(param).var`之后展开预先定义的变量名

```go
package io

func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return
}
```

```go
package utils

func GetContext(param []string) (ctx contex.Context, err error) {
    ctx = context.TODO()
    
    err = FuncA()
    if err != nil {
        return
    }
    
    ctx, err = FuncB()
    return
}
```

### 更多格式之外的代码规范及注意事项参考

* [Effective Go](https://golang.org/doc/effective_go.html#named-results)
* [Go standards and style guidelines](https://docs.gitlab.com/ee/development/go_guide/)
* [Uber](https://github.com/xxjwxc/uber_go_guide_cn)

## 目录说明

### `/api`

模板文件，第三方库所需的数据文件，JSON，proto定义等

### `/cmd`

各个项目模块的[`main`]函数入口程序

### `/config`

项目配置文件，监听端口，数据库信息等

### `/githook`

githook脚本文件

### `/init`

读取配置文件到全局变量，初始化项目信息

### `/internal`

项目内部引用的包

    /app 对应cmd目录，只被对应模块引用
    /pkg 模块公用代码
    
### `/tools`

独立于业务之外的工具函数


