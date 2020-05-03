package loads

import (
	"main/instructions/base"
	"main/rtda"
)

type ALOAD struct{ base.Index8Instruction }
type ALOAD_0 struct{ base.NoOperandsInstruction }
type ALOAD_1 struct{ base.NoOperandsInstruction }
type ALOAD_2 struct{ base.NoOperandsInstruction }
type ALOAD_3 struct{ base.NoOperandsInstruction }

func _aload(frame *rtda.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index) // 根据索引index从局部变量表获取值
	frame.OperandStack().PushRef(ref)      // 推入操作数栈
}

func (self *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, self.Index)
}

func (self *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

func (self *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

func (self *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

func (self *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}

