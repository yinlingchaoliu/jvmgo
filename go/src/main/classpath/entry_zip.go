package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// 读取zip文件类
type ZipEntry struct {
	absPath string // 存放zip或jar文件的目录（绝对路径）
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 1. 打开zip文件
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()

	// 2. 遍历zip或jar中的所有文件，找到要搜索的className文件，读取为[]byte
	for _, f := range r.File {
		//寻找指定类文件
		if f.Name == className {
			//打开文件
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()

			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, err
		}
	}

	return nil, nil, errors.New("class not found:" + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}
