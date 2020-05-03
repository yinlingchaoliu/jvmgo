package references

import (
	"main/instructions/base"
	"main/rtda"
	"main/rtda/heap"
)

// new: 创建类实例
type NEW struct {
	base.Index16Instruction // 用于从常量池寻找类符号引用
}

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()            // 1. 获取当前栈帧所在类的常量池
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef) // 2. 获取类符号引用
	class := classRef.ResolvedClass()                      // 3. 根据类符号引用创建类
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()          // 4. 创建对象
	frame.OperandStack().PushRef(ref) // 5. 将引用对象push到栈顶
}
