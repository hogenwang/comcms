package common

import (
	"os"
	"path/filepath"
	"strings"
)

//获取目录下所有文件,返回一个切片
func GetFolderFiles(folder string, allowFileType string) ([]string, error) {
	list := make([]string, 0)
	if allowFileType == "" {
		allowFileType = "gif|jpeg|jpg|png|bmp|html|htm|tpl|css"
	}
	err := filepath.Walk(folder, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return nil
		}
		if f.IsDir() {
			return nil
		}
		ext := path[strings.LastIndex(path, ".")+1:]
		if strings.Contains(allowFileType, ext) {
			file := f.Name()
			list = append(list, file)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return list, nil
}
