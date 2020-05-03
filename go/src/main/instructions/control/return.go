package control

import (
	"main/instructions/base"
	"main/rtda"
)

// 返回 Void
type RETURN struct {
	base.NoOperandsInstruction
}

// 返回 引用
type ARETURN struct {
	base.NoOperandsInstruction
}

// 返回 int
type IRETURN struct {
	base.NoOperandsInstruction
}

// 返回 long
type LRETURN struct {
	base.NoOperandsInstruction
}

// 返回 float
type FRETURN struct {
	base.NoOperandsInstruction
}

// 返回 double
type DRETURN struct {
	base.NoOperandsInstruction
}

func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame() // 直接弹出当前栈帧
}

// 将当前栈帧的返回值（操作数栈顶）移除并推入调用者栈帧的操作数栈顶
// 将当前栈帧从线程中移除
func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame() // 当前方法的栈帧
	invokerFrame := thread.TopFrame() // 调用当前方法的前一个方法的栈帧

	retVal := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(retVal)
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame() // 当前方法的栈帧
	invokerFrame := thread.TopFrame() // 调用当前方法的前一个方法的栈帧

	retVal := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(retVal)
}

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame() // 当前方法的栈帧
	invokerFrame := thread.TopFrame() // 调用当前方法的前一个方法的栈帧

	retVal := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(retVal)
}

func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame() // 当前方法的栈帧
	invokerFrame := thread.TopFrame() // 调用当前方法的前一个方法的栈帧

	retVal := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(retVal)
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame() // 当前方法的栈帧
	invokerFrame := thread.TopFrame() // 调用当前方法的前一个方法的栈帧

	retVal := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(retVal)
}
