package comparisons

import (
	"main/instructions/base"
	"main/rtda"
)

/*
比较指令：两个类 - 一类将比较结果push到操作数栈顶；一类根据比较结果跳转
if<cond>
将操作数栈顶的int变量弹出，之后与0比较，满足条件就跳转
*/
type IFEQ struct{ base.BranchInstruction } // ==0
type IFNE struct{ base.BranchInstruction } // !=0
type IFLT struct{ base.BranchInstruction } // <0
type IFLE struct{ base.BranchInstruction } // <=0
type IFGT struct{ base.BranchInstruction } // >0
type IFGE struct{ base.BranchInstruction } // >=0

func (self *IFEQ) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v==0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNE) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v!=0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFLT) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v<0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFLE) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v<=0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFGT) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v>0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFGE) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v>=0 {
		base.Branch(frame, self.Offset)
	}
}
