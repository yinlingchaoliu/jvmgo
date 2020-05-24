package rtda

import "main/rtda/heap"

type Thread struct {
	pc    int    // 程序计数器 -> 虚拟机指令质地
	stack *Stack //java虚拟机栈
}

func NewTread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

// 栈帧入栈
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

// 栈帧出栈
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

// 获取当前栈帧
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

//获得栈顶
func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}

//获得寄存器
func (self *Thread) PC() int {
	return self.pc
}

//设置寄存器
// Setter
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

//判断栈是否为空
func (self *Thread) IsStackEmpty() bool {
	return self.stack.IsEmpty()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}

//测试帧
func (self *Thread) NewTestFrame(maxLocals, maxStack uint) *Frame {
	return newTestFrame(self, maxLocals, maxStack)
}

//获得当前线程所有栈 todo exception
func (self *Thread) GetFrames() []*Frame {
	return self.stack.getFrames()
}

func (self *Thread) ClearStack() {
	self.stack.clear()
}