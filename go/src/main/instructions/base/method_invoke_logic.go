package base

import (
	"fmt"
	"main/rtda"
	"main/rtda/heap"
)

// invokerFrame - 调用当前的方法的哪一个方法栈帧
// method - 当前方法，即被 invokerFrame 调用的方法
func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	// 1.使用同一个线程为当前方法创建栈帧并压入线程栈顶
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	// 2. 获取当前方法需要的参数个数，并从调用者 invokerFrame 的操作数栈中弹出制定个数个参数，放到当前方法的栈帧的本地变量中
	argSlotCount := int(method.ArgSlotCount())
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

	//hack
	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method: %v.%v%v\n",
				method.Class().Name(), method.Name(), method.Descriptor()))
		}
	}
}
