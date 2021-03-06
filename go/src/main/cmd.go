package main

import "flag"
import "fmt"
import "os"

//命令行
// java [-options] class [args...]
//todo 命令行参数
type Cmd struct {
	helpFlag         bool //java -help
	versionFlag      bool //java -version
	verboseClassFlag bool  //类加载信息输出
	verboseInstFlag  bool  //指令信息输出
	cpOption         string
	XjreOption       string   // 指定jre启动类的目录
	testOption       string   // 指定测试方法
	class            string   // 执行主类
	args             []string // 附加参数
}

//将flag参数转成cmd
func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose:class", false, "print classloader info")
	flag.BoolVar(&cmd.verboseInstFlag, "verbose:inst", false, "print inst info")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	//增加测试方法
	flag.StringVar(&cmd.testOption, "test", "", "test")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	//parse失败 会执行 printUsage
	flag.Parse()

	//解析剩余参数
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

//使用范例
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}