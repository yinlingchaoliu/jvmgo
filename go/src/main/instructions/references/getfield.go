package references

import (
	"main/instructions/base"
	"main/rtda"
	"main/rtda/heap"
)

// 获取实例变量值
type GET_FIELD struct {
	base.Index16Instruction // 用于从常量池寻找字段符号引用
}

func (self *GET_FIELD) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	// 1. 获取字段符号引用
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	// 2. 将字段符号引用解析为Field
	field := fieldRef.ResolvedField()
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	stack := frame.OperandStack()
	// 3. 获取对象引用
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	slots := ref.Fields()
	// 4. 从对象引用的实例变量列表中获取值，push到操作数栈
	switch field.Descriptor()[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(field.SlotId()))
	case 'F':
		stack.PushFloat(slots.GetFloat(field.SlotId()))
	case 'J':
		stack.PushLong(slots.GetLong(field.SlotId()))
	case 'D':
		stack.PushDouble(slots.GetDouble(field.SlotId()))
	case 'L', '[': // 对象或数组
		stack.PushRef(slots.GetRef(field.SlotId()))
	}
}
