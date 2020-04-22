package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry //启动类搜索
	extClasspath  Entry //扩展类搜索
	userClasspath Entry //用户类搜索
}

//创建解析器
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	//解析启动类加载
	cp.parseBootAndExtClasspath(jreOption)
	//解析用户类加载
	cp.parseUserClasspath(cpOption)
	return cp
}

//如果用户没有传入 -cp 选项 则使用当前目录
// 读取文件名称为className的class文件
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	// 1. 从启动类路径寻找读取 <className>.class 类
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	// 2. 从扩展类路径寻找读取 <className>.class 类
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	// 3. 从用户类路径寻找读取 <className>.class 类
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	self.bootClasspath = newWildcardEntry(filepath.Join(jreDir, "lib", "*"))       // jre/lib/*
	self.extClasspath = newWildcardEntry(filepath.Join(jreDir, "lib", "ext", "*")) // jre/lib/ext/*
}

//获得jre目录
func getJreDir(jreOption string) string {
	// 先读取命令行参数-Xjre，如果存在，直接返回（为了简化，不做错误输入的处理）
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	// 如果命令行没有传入-Xjre，使用JAVA_HOME/jre
	if javaHome := os.Getenv("JAVA_HOME"); javaHome != "" {
		return filepath.Join(javaHome, "jre")
	}
	panic("Can't find jre folder")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}
