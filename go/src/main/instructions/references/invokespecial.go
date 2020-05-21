package references

import (
	"main/instructions/base"
	"main/rtda"
	"main/rtda/heap"
)

// 调用实例方法:私有方法 和 构造函数
type INVOKE_SPECIAL struct{ base.Index16Instruction }

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedClass := methodRef.ResolvedClass()
	resolveMethod := methodRef.ResolveMethod()

	if resolveMethod.Name() == "<init>" && resolveMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}

	if resolveMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// todo  resolveMethod.ArgSlotCount()-1 避免数组越界
	ref := frame.OperandStack().GetRefFromTop(resolveMethod.ArgSlotCount()-1) // 弹出 this 引用
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	base.InvokeMethod(frame, resolveMethod)
}
