package references

import (
	"fmt"
	"main/instructions/base"
	"main/rtda"
	"main/rtda/heap"
)

// 实例方法
type INVOKE_VIRTUAL struct{ base.Index16Instruction }

// hack
func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolveMethod := methodRef.ResolveMethod()

	if resolveMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	ref := frame.OperandStack().GetRefFromTop(resolveMethod.ArgSlotCount() - 1)

	if ref == nil {
		// hack!
		if methodRef.Name() == "println" {
			_println(frame.OperandStack(), methodRef.Descriptor())
			return
		}
		panic("java.lang.NullPointerException")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil && methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}

func _println(stack *rtda.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V":
		fmt.Printf("[hack_println: %v ]\n", stack.PopInt() != 0)
	case "(C)V":
		fmt.Printf("[hack_println: %c ]\n", stack.PopInt())
	case "(B)V":
		fmt.Printf("[hack_println: %v ]\n", stack.PopInt())
	case "(S)V":
		fmt.Printf("[hack_println: %v ]\n", stack.PopInt())
	case "(I)V":
		fmt.Printf("[hack_println: %v ]\n", stack.PopInt())
	case "(J)V":
		fmt.Printf("[hack_println: %v ]\n", stack.PopLong())
	case "(F)V":
		fmt.Printf("[hack_println: %v ]\n", stack.PopFloat())
	case "(D)V":
		fmt.Printf("[hack_println: %v ]\n", stack.PopDouble())
	case "(Ljava/lang/String;)V":
		jStr := stack.PopRef()
		//将内存字符串utf-16 转化为utf-8
		goStr := heap.GoString(jStr)
		fmt.Printf("[hack_println: %s ]\n", goStr)
	default:
		panic("[hack_println: " + descriptor+ " ]")
	}
	stack.PopRef()
}
