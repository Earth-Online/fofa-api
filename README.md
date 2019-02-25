# fofa-api : fofa vip api

## Overview [![GoDoc](https://godoc.org/github.com/Earth-Online/fofa-api?status.svg)](https://godoc.org/github.com/Earth-Online/fofa-api/client) [![Build Status](https://travis-ci.org/Earth-Online/fofa-api.svg?branch=master)](https://travis-ci.org/Earth-Online/fofa-api)

提供了fofa会员的api,如列出自己所有poc,下载poc文件, 提交poc,搜索poc

## Install

```
go get github.com/Earth-Online/fofa-api
```

### 例子
```go
package main

import (
	  "fmt"
      "github.com/Earth-Online/fofa-api/client"
)


func main(){
	var email = "test@qq.com"
    var token = "123456"
    user := client.NewUser(email, token)
    err := user.Me()
    if err != nil{
      panic(err)  
    } 
    fmt.Print(user.Email)
    fmt.Print(user.UserName)
    resp, err := user.GetPoc()
    fmt.Print(resp)
}
```
## License
Apache 2.0.
