# Docs



## 环境 environment

```sh
$ go version
go version go1.18 darwin/amd64
```

## 如何创建一个新项目 create a new project


(1) 新建目录  Create a directory for Go module source code

```
$ mkdir gotools
$ cd gotools
```

(2) 初始化 Start module using the go mod init command

```
$ go mod init github.com/wangyongtao/gotools
```

(3) 写一些代码 write some code :

```
$ vi main.go
```

main.go: 

```go
package main

import (
	"log"
)

func main() {
	log.Println("hello, my go tools")
}
```

(4) 编译并执行代码 compile and run: 

```
$ go run main.go
```

test: 

$ go test -v util/password_test.go util/password.go


(5) Create a new repository in github.com

(6) Git初始化 git init && git remote add: 

```sh
$ cd gotools
$ git init 
$ git remote add origin git@github.com:wangyongtao/gotools.git
```

(7) 代码推送 git push:

```sh
$ git status
$ git add .
$ git commit -m "init"
$ git push --set-upstream origin master
```

(8) 发一个版本 Draft a new release in github.com


(9) 在其他项目中调用本库  Call your code from another module

```sh
$ mkdir hello
$ cd hello
$ go mod init example.com/hello
$ go get -u github.com/wangyongtao/gotools
```

## 参考 Reference

https://go.dev/doc  
https://go.dev/doc/tutorial/create-module  


·END·