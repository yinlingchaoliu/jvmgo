package main

import "flag"
import "fmt"
import "os"

// java [-options] class [args...]
type Cmd struct {
	helpFlag    bool   //java -help
	versionFlag bool   //java -version
	cpOption    string
	class       string // 执行主类
	args        []string // 附加参数
}

//将flag参数转成cmd
func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
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