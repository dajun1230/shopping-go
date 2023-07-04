package utils

import (
	"errors"
	"os"
)

/**
 * @params path 文件名称
 * 文件目录是否存在
 */
func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}
