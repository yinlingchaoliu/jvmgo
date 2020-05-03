package stack

import (
	"main/instructions/base"
	"main/rtda"
)

/*
dup指令：复制操作数栈的栈顶变量，然后推入操作数栈
1. dup
2. dup_x1
3. dup_x2
4. dup2
5. dup2_x1
6. dup2_x2
*/

type DUP struct{ base.NoOperandsInstruction }

func (self *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}
