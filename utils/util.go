// Package utils 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-28 20:54
// @description:
package utils

import "os"

// MkdirIfNotExist makes directories if the input path is not exists
func MkdirIfNotExist(dir string) error {
	if len(dir) == 0 {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}

	return nil
}
