package rtda

import "main/rtda/heap"

//todo 帧计算 要重新重点看
// 栈帧 stack frame
// 栈帧
type Frame struct {
	lower        *Frame        // 下一个栈帧
	localVars    LocalVars     // 本地变量表
	operandStack *OperandStack // 操作数栈
	thread       *Thread       // 栈帧所属的线程
	nextPC       int
	method       *heap.Method // 栈帧所在的方法
}

// 执行方法所需的局部变量大小 maxLocals 和操作数栈深度会由编译器预先计算好，
// 存储在 class 文件的 method_info 方法表结构中的 code 属性中
func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

//测试工具
func newTestFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}

// Setter
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}

// Getter
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) Method() *heap.Method {
	return self.method
}

