package main

import (
	"fmt"
	"main/classpath"
	"main/rtda/heap"
	"strings"
)

/**
 * 模仿虚拟机启动整个流程
 * @author chentong
 * 2020-05-05
 */
func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else if cmd.testOption == "cmd" {
		parseCmdLine(cmd)
	} else if cmd.testOption == "classpath" {
		parseClasspath(cmd)
	} else if cmd.testOption == "classfile" {
		parseClassFile(cmd)
	} else if cmd.testOption == "rtda" {
		parseRtda(cmd)
	} else if cmd.testOption == "interpret" {
		parseInterpret(cmd)
	} else if cmd.testOption == "classloader" {
		parseClassLoader(cmd)
	} else if cmd.testOption == "return" {
		parseReturn(cmd)
	} else if cmd.testOption == "array" {
		parseArray(cmd)
	} else if cmd.testOption == "string" {
		parseStringArgs(cmd)
	} else {
		startJvm(cmd)
	}
}

//启动jvm
func startJvm(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	//获得classLoader
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	//获得加载类名字
	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	//获得main方法
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		//增加命令行参数
		Interpret(mainMethod, cmd.verboseInstFlag,cmd.args)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
