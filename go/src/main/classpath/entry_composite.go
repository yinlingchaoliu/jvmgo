package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

//组合式搜索类文件
func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	//对CompositeEntry 进行遍历 执行readClass方法
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		//查找成功 className
		if err == nil {
			return data, from, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

//将className 拼成一串
func (self CompositeEntry) String() string{
	//创建字符串数组
	strs := make([]string,len(self))

	for i,entry :=range self{
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparator)
}