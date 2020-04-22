package classpath

import "os"
import "strings"

//分隔符 ":"
const pathListSeparator = string(os.PathListSeparator)

//定义接口
type Entry interface {
	// 寻找和读取 class 文件
	// 入参：className - class文件的相对路径，eg. 如果要读取 java.lang.Object 类，则className = java/lang/Object.class
	// 返回值：
	// 1. 读取到的class文件内容的[]byte
	// 2. 最终定位到包含className文件的Entry对象
	// 3. 错误信息error
	readClass(className string) ([]byte, Entry, error)
	//获得className
	String() string
}

//根据参数类型创建不同搜索模式
func newEntry(path string) Entry {

	//读取多个className文件 java -cp path1/classes:path2/classes
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	// 读取path下所有jar文件的className文件  java -cp path/*
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	// 从path/lib1.jar下查找并读取className文件：java -cp path/lib1.jar 或者 java -cp path/lib1.zip
	//读取zip/jar下 className文件 :java -cp path/lib1.jar
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".zip") {
		return newZipEntry(path)
	}

	//遍历目录
	return newDirEntry(path)
}
