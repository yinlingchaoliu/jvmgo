package base

import (
	"main/rtda"
	"main/rtda/heap"
)

// jvms 5.5
func InitClass(thread *rtda.Thread, class *heap.Class) {
	class.StartInit()             // 设置开始初始化标志
	scheduleClinit(thread, class) // 准备执行类初始化方法，push到栈顶，下一次执行的时候就会执行<cinit>
	initSuperClass(thread, class) // 循环初始化父类
}

func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		// exec <clinit>
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}
