package util

import (
    "io"
    "os"
    "path/filepath"
    "syscall"
)

// FileExists: Checks whether a file or directory exists
// 如果由 filename 指定的文件或目录 存在则返回 true，否则返回 false
func FileExists(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil || os.IsExist(err)
}

// IsFile: Tells whether the filename is a regular file
// 判断给定文件名是否为一个文件: 是文件返回true
func IsFile(filename string) bool {
    fileInfo, err := os.Stat(filename)
    if err != nil {
        return false
    }

    return !fileInfo.IsDir()
}

// IsDir: Tells whether the filename is a directory
// 判断给定文件名是否是一个目录: 是目录返回true
func IsDir(filename string) bool {
    fileInfo, err := os.Stat(filename)
    if err != nil {
        return false
    }

    return fileInfo.IsDir()
}

// FileSize filesize(): Gets file size
func FileSize(filename string) (int64, error) {
    info, err := os.Stat(filename)
    if err != nil && os.IsNotExist(err) {
        return 0, err
    }

    return info.Size(), nil
}

// CopyFile copy()
func CopyFile(source string, dest string) (bool, error) {
    fd1, err := os.Open(source)
    if err != nil {
        return false, err
    }
    defer fd1.Close()
    fd2, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return false, err
    }
    defer fd2.Close()
    _, e := io.Copy(fd2, fd1)
    if e != nil {
        return false, e
    }
    
    return true, nil
}

// IsReadable is_readable()
func IsReadable(filename string) bool {
    _, err := syscall.Open(filename, syscall.O_RDONLY, 0)

    return err == nil
}

// IsWriteable is_writeable()
func IsWriteable(filename string) bool {
    _, err := syscall.Open(filename, syscall.O_WRONLY, 0)

    return err == nil
}

// RenameFile rename(): Renames a file or directory
func RenameFile(oldName, newName string) error {
    return os.Rename(oldName, newName)
}

// FileMTime filemtime(): Gets file modification time
func FileMTime(filename string) (int64, error) {
    fd, err := os.Open(filename)
    if err != nil {
        return 0, err
    }
    defer fd.Close()
    fileinfo, err := fd.Stat()
    if err != nil {
        return 0, err
    }

    return fileinfo.ModTime().Unix(), nil
}

// UnlinkFile unlink() : Deletes a file
func UnlinkFile(filename string) error {
    return os.Remove(filename)
}

func FileExt(filename string) string {
    return filepath.Ext(filename)
}
