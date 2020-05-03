package main

import (
	"fmt"
	"main/classfile"
	"main/instructions"
	"main/instructions/base"
	"main/rtda"
)

//解释器 外部不能访问 私有方法
func interpret(methodInfo *classfile.MemberInfo){

	//获得method类 code属性
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	thread := rtda.NewTread()
	frame  := thread.NewTestFrame(maxLocals,maxStack)

	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, bytecode)
}

//异常处理 因没有实现return指令 catch异常
func catchErr(frame *rtda.Frame){
	if r:=recover();r!=nil{
		fmt.Printf("LocalVars:%v\n",frame.LocalVars())
		fmt.Printf("OperandStack:%v\n",frame.OperandStack())
		fmt.Printf("no return fun \n")
		//panic(r)
	}
}

//loop执行循环所有方法
func loop(thread *rtda.Thread, bytecode []byte){

	frame:=thread.PopFrame()
	reader:= &base.ByteCodeReader{}

	for{
		//寻找下一个函数 计算pc
		pc:= frame.NextPC()
		thread.SetPC(pc)

		//设置初始值   解码指令
		reader.Reset(bytecode,pc)
		//读取指令集
		opcode:=reader.ReadUint8()
		//指令集转义
		inst:=instructions.NewInstruction(opcode)
		//读取变量
		inst.FetchOperands(reader)
		//获得下一个指令集便宜
		frame.SetNextPC(reader.PC())

		//excute   执行
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}

}

