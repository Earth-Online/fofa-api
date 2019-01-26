### fofa会员api golang 版
提供了fofa会员的api,如列出自己所有poc,下载poc文件

### 例子
```go
package main

var email = "test@qq.com"
var token = "123456"

def main(){
    user := NewUser(email, token)
    err := user.Me()
	  if err != nil{
      panic(err)  
    } 
    fmt.Print(user.Email)
    fmt.Print(user.Username)
    resp, err := user.GetPoc()
    fmt.Print(resp)
}
```
