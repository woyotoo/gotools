# My Go Tools (Golang实用工具)


## 环境 environment

$ go version
go version go1.18 darwin/amd64


## init


$ mkdir gotools
$ cd gotools

$ go mod init github.com/wangyongtao/gotools

$ vi main.go

```go
package main

import (
	"log"
)

func main() {
	log.Println("hello, my go tools")
}
```

$ go run main.go


go test -v util/password_test.go util/password.go


$ cd gotools
$ git init 
$ git status
$ git remote add origin git@github.com:wangyongtao/gotools.git


## 参考 Reference

