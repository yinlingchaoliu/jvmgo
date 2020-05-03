package comparisons

import (
	"main/instructions/base"
	"main/rtda"
)

/*
比较指令：两个类 - 一类将比较结果push到操作数栈顶；一类根据比较结果跳转
if_icmp<cond>
将操作数栈顶的两个int变量弹出，之后比较，满足条件就跳转
*/
type IF_ICMPEQ struct{ base.BranchInstruction } // ==
type IF_ICMPNE struct{ base.BranchInstruction } // !=
type IF_ICMPLT struct{ base.BranchInstruction } // <
type IF_ICMPLE struct{ base.BranchInstruction } // <=
type IF_ICMPGT struct{ base.BranchInstruction } // >
type IF_ICMPGE struct{ base.BranchInstruction } // >=

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 == v2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 != v2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPLT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 < v2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPLE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 <= v2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPGT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 > v2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 >= v2 {
		base.Branch(frame, self.Offset)
	}
}
