package utils

import (
	"log"
	"os"
	"path/filepath"
)

// 判断目录是否存在
func PathExists(path string) (bool, error) {
	// 获取文件信息
	fi, err := os.Stat(path)
	// 如果没有错误
	if err == nil {
		// 判断是否是目录 返回true
		return fi.IsDir(), nil
	}
	// 如果错误是不存在
	if os.IsNotExist(err) {
		// 返回false
		return false, nil
	}
	// 否则返回错误
	return false, err
}

// 判断文件是否存在
func FileExists(path string) (bool, error) {
	// 获取文件信息
	fi, err := os.Stat(path)
	// 如果没有错误
	if err == nil {
		// 判断是否是目录 返回true
		return !fi.IsDir(), nil
	}
	// 如果错误是不存在
	if os.IsNotExist(err) {
		// 返回false
		return false, nil
	}
	// 否则返回错误
	return false, err
}

// 判断文件或目录是否存在
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	//  err == nil 就是存在
	if err == nil {
		return true, nil
	}
	// err != nil 就是不一定是不存在，需要errors.Is(err, fs.ErrNotExist).判断
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetProgramPath 获取程序所在目录
// @return string 文件夹路径
func GetProgramPath() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(exePath)
}
