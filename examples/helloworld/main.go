package main

import (
    "fmt"

    // 引入工具库
    "github.com/wangyongtao/gotools/util"
)

// 演示加密字符串、校验加密结果

func main() {
    fmt.Println("hello, world!")

    // 加密：encrypt
    // 用户密码，需要加密存储
    pwd := "admin"
    hash, _ := util.PasswordHash(pwd)

    fmt.Println("==> 密码加密:")
    fmt.Println("--> 输入的密码:", pwd)
    fmt.Println("--> 生成的hash:", hash)
    // hash: $2a$10$tL00vjlVzidXygcmmW7naObKkI2CgoH9lcaT/DYAFWVAbVO/PVMiu

    // 校验 verify
    // 登录账户时，将用户输入的密码加密后，与数据中查询的加密密码对比，校验通过则说明用户密码输入正确 
    fmt.Println("==> 对比加密结果:")
    match := util.PasswordVerify(pwd, hash)

    fmt.Println("--> 验证结果:", match)
}