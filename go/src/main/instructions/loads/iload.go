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
type ILOAD struct{ base.Index8Instruction }
type ILOAD_0 struct{ base.NoOperandsInstruction }
type ILOAD_1 struct{ base.NoOperandsInstruction }
type ILOAD_2 struct{ base.NoOperandsInstruction }
type ILOAD_3 struct{ base.NoOperandsInstruction }

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index) // 根据索引index从局部变量表获取值
	frame.OperandStack().PushInt(val)      // 推入操作数栈
}

func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, self.Index)
}

func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
