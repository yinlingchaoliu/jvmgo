package references

import (
	"main/instructions/base"
	"main/rtda"
	"main/rtda/heap"
)

// checkcast: 判断对象是否是某个类的实例 或者 对象的类是否实现了某个接口
// 与 instanceof 类似，但是不操作操作数栈
type CHECK_CAST struct {
	base.Index16Instruction // 用于从常量池寻找类符号引用
}

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
	// 1. 从操作数栈获取对象引用ref
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	// (Integer)null -> null引用可以转换为任何类型
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()            // 2. 获取当前栈帧所在类的常量池
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef) // 3. 获取类符号引用
	class := classRef.ResolvedClass()                      // 4. 根据类符号引用创建类
	if !ref.IsInstanceOf(class) { // 5. 判断 ref instanceof class
		panic("java.lang.ClassCastException") // 6. 如果不是，抛异常
	}
}
