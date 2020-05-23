package constants

import (
	"main/instructions/base"
	"main/rtda"
	"main/rtda/heap"
)

type LDC struct {
	base.Index8Instruction
}
type LDC_W struct {
	base.Index16Instruction
}
type LDC2_W struct {
	base.Index16Instruction
}

func (self *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

func (self *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

func (self *LDC2_W) Execute(frame *rtda.Frame) {
	// 1. 从运行时常量池获取常量c
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.Index)

	// 2. 将常量c压入操作数栈
	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}

func _ldc(frame *rtda.Frame, index uint) {
	// 1. 从运行时常量池获取常量c
	stack := frame.OperandStack()
	class := frame.Method().Class()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)

	// 2. 将常量c压入操作数栈
	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	case string:  //todo 支持字符串 压栈
		internedStr := heap.JString(class.Loader(),c.(string))
		stack.PushRef(internedStr)
	case *heap.ClassRef:
		classRef := c.(*heap.ClassRef)
		classObj := classRef.ResolvedClass().JClass()
		stack.PushRef(classObj)
	default:
		panic("todo:ldc!")
	}
}
