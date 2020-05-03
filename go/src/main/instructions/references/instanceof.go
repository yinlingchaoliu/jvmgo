package references

import (
	"main/instructions/base"
	"main/rtda"
	"main/rtda/heap"
)

// instanceof: 判断对象是否是某个类的实例 或者 对象的类是否实现了某个接口
type INSTANCE_OF struct {
	base.Index16Instruction // 用于从常量池寻找类符号引用
}

func (self *INSTANCE_OF) Execute(frame *rtda.Frame) {
	// 1. 从操作数栈获取对象引用ref
	stack := frame.OperandStack()
	ref := stack.PopRef()
	// null instanceof Xxx -> false
	if ref == nil {
		stack.PushInt(0)
		return
	}

	cp := frame.Method().Class().ConstantPool()            // 2. 获取当前栈帧所在类的常量池
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef) // 3. 获取类符号引用
	class := classRef.ResolvedClass()                      // 4. 根据类符号引用创建类
	if ref.IsInstanceOf(class) { // 5. 判断 ref instanceof class
		stack.PushInt(1) // 6. 将结果压入栈
	} else {
		stack.PushInt(0)
	}
}
