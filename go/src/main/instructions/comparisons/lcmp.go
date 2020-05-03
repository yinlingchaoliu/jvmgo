package comparisons

import (
	"main/instructions/base"
	"main/rtda"
)

/*
比较指令：两个类 - 一类将比较结果push到操作数栈顶；一类根据比较结果跳转
*/
type LCMP struct{ base.NoOperandsInstruction }

func (self *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
