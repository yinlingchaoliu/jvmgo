package math

import (
	"main/instructions/base"
	"main/rtda"
	"math"
)

/*
rem求余指令：从操作数栈的栈顶获取两个变量，之后进行求余操作（栈顶元素为分母），计算之后再push到操作数栈中
*/

type IREM struct{ base.NoOperandsInstruction }
type LREM struct{ base.NoOperandsInstruction }
type DREM struct{ base.NoOperandsInstruction }
type FREM struct{ base.NoOperandsInstruction }

func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}

func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushLong(result)
}

func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2))) // 浮点型具有 Infinity（无穷大）值，如果除0，也不会抛出ArithmeticException异常
	stack.PushFloat(result)
}

func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2) // 浮点型具有 Infinity（无穷大）值，如果除0，也不会抛出ArithmeticException异常
	stack.PushDouble(result)
}
