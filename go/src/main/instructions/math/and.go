package math

import (
	"main/instructions/base"
	"main/rtda"
)

/*
布尔指令：从操作数栈的栈顶获取两个变量，之后进行三种之一的操作操作（以按位与为例），计算之后，再push到栈中
只能操作int和long变量，分为按位与（and）、按位或（or）。按位异或（xor）
*/

type IAND struct{ base.NoOperandsInstruction }
type LAND struct{ base.NoOperandsInstruction }

func (self *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

func (self *LAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
