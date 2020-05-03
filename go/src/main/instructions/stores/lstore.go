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
type LSTORE struct{ base.Index8Instruction }
type LSTORE_0 struct{ base.NoOperandsInstruction }
type LSTORE_1 struct{ base.NoOperandsInstruction }
type LSTORE_2 struct{ base.NoOperandsInstruction }
type LSTORE_3 struct{ base.NoOperandsInstruction }

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong() // 从操作数栈pop
	frame.LocalVars().SetLong(index, val) // 放到局部变量表中指定的索引index位置处

}

func (self *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, self.Index)
}

func (self *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}

func (self *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}

func (self *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}

func (self *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}
