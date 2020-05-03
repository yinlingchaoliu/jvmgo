package references

import (
	"main/instructions/base"
	"main/rtda"
	"main/rtda/heap"
)

// 设置实例变量值
type PUT_FIELD struct {
	base.Index16Instruction // 用于从常量池寻找字段符号引用
}

func (self *PUT_FIELD) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	// 1. 获取字段符号引用
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	// 2. 将字段符号引用解析为Field
	field := fieldRef.ResolvedField()
	class := field.Class()

	// 不是实例变量
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// final 只能在构造器初始化方法中进行赋值
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	stack := frame.OperandStack()
	switch field.Descriptor()[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt() // 3. 获取将要设置的属性值
		ref := stack.PopRef() // 4. 获取将要设置的值所属的对象引用
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetInt(field.SlotId(), val) // 5. 将属性值设置到属性中
	case 'F':
		val := stack.PopFloat() // 3. 获取将要设置的属性值
		ref := stack.PopRef() // 4. 获取将要设置的值所属的对象引用
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetFloat(field.SlotId(), val) // 5. 将属性值设置到属性中
	case 'J':
		val := stack.PopLong() // 3. 获取将要设置的属性值
		ref := stack.PopRef() // 4. 获取将要设置的值所属的对象引用
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetLong(field.SlotId(), val) // 5. 将属性值设置到属性中
	case 'D':
		val := stack.PopDouble() // 3. 获取将要设置的属性值
		ref := stack.PopRef() // 4. 获取将要设置的值所属的对象引用
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetDouble(field.SlotId(), val) // 5. 将属性值设置到属性中
	case 'L', '[': // 对象或数组
		val := stack.PopRef() // 3. 获取将要设置的属性值
		ref := stack.PopRef() // 4. 获取将要设置的值所属的对象引用
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetRef(field.SlotId(), val) // 5. 将属性值设置到属性中
	}
}

