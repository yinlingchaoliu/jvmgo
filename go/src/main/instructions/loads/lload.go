package loads

import (
	"main/instructions/base"
	"main/rtda"
)

/*
load指令：
1. aload - 引用类型
2. lload - long
3. fload - float
4. dload - double
5. xload - 数组
6. iload - int
以 iload 为例，将 int 类型的数据根据索引index从局部变量表中查出
然后推入操作数栈
 */
type LLOAD struct{ base.Index8Instruction }
type LLOAD_0 struct{ base.NoOperandsInstruction }
type LLOAD_1 struct{ base.NoOperandsInstruction }
type LLOAD_2 struct{ base.NoOperandsInstruction }
type LLOAD_3 struct{ base.NoOperandsInstruction }

func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index) // 根据索引index从局部变量表获取值
	frame.OperandStack().PushLong(val)      // 推入操作数栈
}

func (self *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, self.Index)
}

func (self *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

func (self *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

func (self *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

func (self *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}
