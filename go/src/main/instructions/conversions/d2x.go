package conversions

import (
	"main/instructions/base"
	"main/rtda"
)

/*
对栈顶元素进行类型转换，之后在push到栈顶：
1. i2x - int变量强转成其他类型
2. l2x - long变量强转成其他类型
3. f2x - float变量强转成其他类型
4. d2x - double变量强转成其他类型
以 d2x 为例
*/
type D2F struct{ base.NoOperandsInstruction }
type D2I struct{ base.NoOperandsInstruction }
type D2L struct{ base.NoOperandsInstruction }

func (self *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

func (self *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int64(d)
	stack.PushLong(i)
}

func (self *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := float32(d)
	stack.PushFloat(i)
}
