package main

import (
	"fmt"
	"main/classpath"
	"main/instructions/base"
	"main/rtda"
	"main/rtda/heap"
	"strings"
)

//定义jvm
type JVM struct {
	cmd         *Cmd              //命令行
	classLoader *heap.ClassLoader //类加载器
	mainThread  *rtda.Thread      //主线程
}

//新建虚拟机
func newJVM(cmd *Cmd) *JVM {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	return &JVM{
		cmd:         cmd,
		classLoader: classLoader,
		mainThread:  rtda.NewTread(),
	}
}

//启动虚拟机
func (self *JVM) start() {
	//暂时未能真正启动VM
	//self.initJVM()
	self.execMain()
}

//初始化虚拟机
func (self *JVM) initJVM() {
	vmClass := self.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(self.mainThread, vmClass)
	interpret(self.mainThread, self.cmd.verboseInstFlag)
}

//运行main主方法
func (self *JVM) execMain() {

	//获得加载类名字
	className := strings.Replace(self.cmd.class, ".", "/", -1)
	mainClass := self.classLoader.LoadClass(className)
	//获得main方法
	mainMethod := mainClass.GetMainMethod()

	if mainMethod == nil {
		//增加命令行参数
		fmt.Printf("Main method not found in class %s\n", self.cmd.class)
		return
	}

	frame := self.mainThread.NewFrame(mainMethod) //创建栈帧

	//字符串参数
	jArgs := self.createArgsArray()
	frame.LocalVars().SetRef(0, jArgs)
	self.mainThread.PushFrame(frame)              //将栈帧push线程stack中

	interpret(self.mainThread,self.cmd.verboseInstFlag)
}

//创建args数组
func (self *JVM)createArgsArray() *heap.Object {
	//加载class类
	stringClass := self.classLoader.LoadClass("java/lang/String")
	argsArr := stringClass.ArrayClass().NewArray(uint(len(self.cmd.args)))
	jArgs := argsArr.Refs()
	for i, arg := range self.cmd.args {
		jArgs[i] = heap.JString(self.classLoader, arg)
	}
	return argsArr
}