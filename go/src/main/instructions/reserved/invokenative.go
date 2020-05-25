package reserved

import (
	"main/instructions/base"
	"main/rtda"
	"main/native"
	_ "main/native/java/lang" //@todo init 注册
	_ "main/native/sun/misc"  //@todo 注册VM
	_ "main/native/sun/reflect"  //@todo 反射
	_"main/native/java/security" //@todo doPrivileged
	_"main/native/java/util/concurrent/atomic" //@todo juc
)

type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	//调用本地方法
	nativeMethod(frame)
}
