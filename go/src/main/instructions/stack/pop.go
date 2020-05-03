package stack

import (
	"main/instructions/base"
	"main/rtda"
)

/*
pop指令：将操作数栈的栈顶变量弹出
1. pop - 只能弹出占用一个操作数栈位置(slot)的类型，例如 int
2. pop2 - 可以弹出占用2个操作数栈位置(slot)的类型，例如 long、double
*/

type POP struct{ base.NoOperandsInstruction }
type POP2 struct{ base.NoOperandsInstruction }

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
