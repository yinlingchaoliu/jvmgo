package main

import (
	"fmt"
	"main/classfile"
	"main/classpath"
	"main/rtda"
	"main/rtda/heap"
	"strings"
)

/**
 *  编写单元测试用例，避免和main函数混在一起
 *  @author chentong
 */
//测试命令行
func parseCmdLine(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.cpOption, cmd.class, cmd.args)
}

//测试classpath
func parseClasspath(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}

	fmt.Printf("class data:%v\n", classData)
}

//测试加载classfile
func parseClassFile(cmd *Cmd) {
	classpath := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classfile := loadClass(className, classpath)
	printClassFile(classfile)
}

//加载ClassFile
func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {

	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	classfile, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return classfile
}

//打印classFile
func printClassFile(cf *classfile.ClassFile) {
	fmt.Printf("magic: 0xCAFEBABE\n")
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}
}

//测试rtda
func parseRtda(cmd *Cmd) {
	//todo 后续修改
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetLong(1, 4343434)
	vars.SetFloat(2, 3.14)
	vars.SetDouble(3, 121.334)
	vars.SetRef(4, nil)
	println(vars.GetInt(0))
	println(vars.GetLong(1))
	println(vars.GetFloat(2))
	println(vars.GetDouble(3))
	println(vars.GetRef(4))
}

func testOperandStack(ops *rtda.OperandStack) {

	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(23232323)
	ops.PushLong(-23232323)
	ops.PushFloat(3.14)
	ops.PushDouble(2.89)
	ops.PushRef(nil)

	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}

//测试解释器和指令集
func parseInterpret(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	//获得classfile
	cf := loadClass(className, cp)
	//获得main函数
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		//解释器执行
		interpretInfo(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}

}

//获得main函数
func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {

	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}

	return nil
}

//测试classloader
func parseClassLoader(cmd *Cmd) {

	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	//获得classLoader  设置打印信息
	classLoader := heap.NewClassLoader(cp, true)
	//获得加载类名字
	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	//获得main方法
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpretMethod(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}

}

//测试函数调用与返回
func parseReturn(cmd *Cmd) {

	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	//获得classLoader
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	//获得加载类名字
	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	//获得main方法
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		Interpret(mainMethod, cmd.verboseInstFlag)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
