package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

//通配符类
//继承自CompositeEntry
func newWildcardEntry(path string) CompositeEntry {
	// 去掉结尾的*号
	baseDir := path[:len(path)-1]

	compositeEntry := []Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {

		if err!=nil{
			return err
		}

		//虚拟机规范: 通配符加载只可以加载baseDir目录下的jar,其子目录下的jar不可以加载
		if info.IsDir() && path !=baseDir {
			return filepath.SkipDir  //跳过
		}

		if strings.HasSuffix(path,".jar"){
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry,jarEntry)
		}

		return nil
	}

	//目录树加载
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
