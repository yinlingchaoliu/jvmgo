package stores

import (
	"main/instructions/base"
	"main/rtda"
)

/*
store指令：(与 load 指令相反)
1. astore - 引用类型
2. lstore - long
3. fstore - float
4. dstore - double
5. xstore - 数组
6. istore - int
以 lstore 为例，将 long 类型的数据先从操作数栈pop，然后放到局部变量表中指定的索引index位置处
 */
type ISTORE struct{ base.Index8Instruction }
type ISTORE_0 struct{ base.NoOperandsInstruction }
type ISTORE_1 struct{ base.NoOperandsInstruction }
type ISTORE_2 struct{ base.NoOperandsInstruction }
type ISTORE_3 struct{ base.NoOperandsInstruction }

func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt() // 从操作数栈pop
	frame.LocalVars().SetInt(index, val) // 放到局部变量表中指定的索引index位置处

}

func (self *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, self.Index)
}

func (self ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}

func (self *ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

func (self *ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

func (self *ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}

