package references

import (
	"main/instructions/base"
	"main/rtda"
	"main/rtda/heap"
)

// 调用静态方法
type INVOKE_STATIC struct{ base.Index16Instruction }

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	method := methodRef.ResolveMethod() // 根据方法符号引用：方法name、方法描述符descriptor、所包含的类指针 查找方法
	if !method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	class := methodRef.ResolvedClass()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	base.InvokeMethod(frame, method)
}
