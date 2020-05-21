package main

import (
	"fmt"
	"main/classfile"
	"main/instructions"
	"main/instructions/base"
	"main/rtda"
	"main/rtda/heap"
)

//正式实现
//解释器
func Interpret(method *heap.Method, logInst bool) {
	thread := rtda.NewTread()        //创建线程
	frame := thread.NewFrame(method) //创建栈帧
	thread.PushFrame(frame)          //将栈帧push线程stack中
	defer catchErr(thread)
	loop(thread, logInst)
}

//执行指令
func loop(thread *rtda.Thread, logInst bool) {

	reader := &base.ByteCodeReader{}

	for {
		//获取当前栈
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		//decode
		reader.Reset(frame.Method().Code(), pc)
		//读取指令opcode
		opcode := reader.ReadUint8()                // 读取操作码 opCode（指令类型）
		inst := instructions.NewInstruction(opcode) // 根据opCode创建相应的指令
		inst.FetchOperands(reader)                  // 从字节码中读取操作数
		frame.SetNextPC(reader.PC())                // 将当前读取到的字节码的位置设置到 frame 的 nextPc 中，用于执行下一条指令

		if logInst {
			logInstruction(frame, inst)
		}

		//执行栈帧
		inst.Execute(frame)

		//线程中栈帧执行完毕退出
		if thread.IsStackEmpty() {
			break
		}
	}

}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

//打印栈信息
func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		//lineNum := method.GetLineNumber(frame.NextPC())
		fmt.Printf("[LogFrame >> pc:%4d %v.%v%v ]\n", frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("[LogInst %v.%v() #%2d %T %v]\n", className, methodName, pc, inst, inst)
}

/*************测试函数***********/
//解释器 外部不能访问 私有方法
func interpretInfo(methodInfo *classfile.MemberInfo) {

	//获得method类 code属性
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	thread := rtda.NewTread()
	frame := thread.NewTestFrame(maxLocals, maxStack)

	thread.PushFrame(frame)

	defer catchFrameErr(frame)
	loopOneMethod(thread, bytecode)
}

//解释器
func interpretMethod(method *heap.Method) {
	thread := rtda.NewTread()        //创建线程
	frame := thread.NewFrame(method) //创建栈帧
	thread.PushFrame(frame)          //将栈帧push线程stack中
	defer catchFrameErr(frame)
	loopOneMethod(thread, method.Code())
}

//异常处理 因没有实现return指令 catch异常
func catchFrameErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		//todo catchErr 异常处理
		fmt.Printf("[catchErr LocalVars:%v]\n", frame.LocalVars())
		fmt.Printf("[catchErr OperandStack:%v]\n", frame.OperandStack())
		fmt.Printf("[catchErr no return fun]\n")
		//panic(r)
	}
}

//loop执行循环所有方法
func loopOneMethod(thread *rtda.Thread, bytecode []byte) {

	frame := thread.PopFrame()
	reader := &base.ByteCodeReader{}

	for {
		//寻找下一个函数 计算pc
		pc := frame.NextPC()
		thread.SetPC(pc)

		//设置初始值   解码指令
		reader.Reset(bytecode, pc)
		//读取指令集
		opcode := reader.ReadUint8()
		//指令集转义
		inst := instructions.NewInstruction(opcode)
		//读取变量
		inst.FetchOperands(reader)
		//获得下一个指令集便宜
		frame.SetNextPC(reader.PC())

		//todo  excute   执行
		fmt.Printf("[loop pc:%2d inst:%T %v]\n", pc, inst, inst)
		inst.Execute(frame)
	}

}