package references

import (
	"main/instructions/base"
	"main/rtda"
	"main/rtda/heap"
)

// 设置static变量值
type PUT_STATIC struct {
	base.Index16Instruction // 用于从常量池寻找字段符号引用
}

func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	// 1. 获取字段符号引用
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	// 2. 将字段符号引用解析为Field
	field := fieldRef.ResolvedField()
	class := field.Class()

	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	// 不是静态变量
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// final 只能在类初始化方法中进行赋值
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	stack := frame.OperandStack()
	// 3. 从操作数栈弹出操作数，放入Field的静态变量中
	switch field.Descriptor()[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		class.StaticVars().SetInt(field.SlotId(), stack.PopInt())
	case 'F':
		class.StaticVars().SetFloat(field.SlotId(), stack.PopFloat())
	case 'J':
		class.StaticVars().SetLong(field.SlotId(), stack.PopLong())
	case 'D':
		class.StaticVars().SetDouble(field.SlotId(), stack.PopDouble())
	case 'L', '[': // 对象或数组
		class.StaticVars().SetRef(field.SlotId(), stack.PopRef())
	}
}
