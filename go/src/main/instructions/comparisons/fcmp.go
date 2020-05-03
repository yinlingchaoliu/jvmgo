package comparisons

import (
	"main/instructions/base"
	"main/rtda"
)

/*
比较指令：两个类 - 一类将比较结果push到操作数栈顶；一类根据比较结果跳转
*/
type FCMPG struct{ base.NoOperandsInstruction } // 如果其中一个数是NaN时，设置为1, G-greater
type FCMPL struct{ base.NoOperandsInstruction } // 如果其中一个数是NaN时，设置为-1, L-less

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func (self *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

func (self *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}
