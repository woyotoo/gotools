package util

import (
    "testing"
)

// cd gotools
// go test -v util/file_test.go util/file.go

// 文件：gotools/util/README.md 
var filename string = "README.md"

// func Test_FileExists(t *testing.T) {

func Test_FileExists(t *testing.T) {
    t.Log("--> Test_FileExists ")

    // filename := "file.go"
    t.Log("--> filename => ", filename)

    // 检查文件是否存在
    exist := FileExists(filename)

    t.Log("--> result exist => ", exist)
    
    // 检查是否为文件
    isFile := IsFile(filename)
    t.Log("--> isFile:", isFile)
    if isFile == false {
        t.Errorf("Test_FileExists IsFile failed. Got false, expected true.")
    }

    // 检查是否为目录
    isDir := IsDir(filename)
    t.Log("--> isDir:", isDir)
    if isDir == true {
        t.Errorf("Test_FileExists IsDir failed. Got false, expected true.")
    }
}