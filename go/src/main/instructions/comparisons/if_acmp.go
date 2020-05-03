package comparisons

import (
	"main/instructions/base"
	"main/rtda"
)

/*
比较指令：两个类 - 一类将比较结果push到操作数栈顶；一类根据比较结果跳转
if_acmp<cond>
if_acmpeq, if_acmpne 将操作数栈顶的两个引用变量弹出，之后比较引用是否相同，满足条件就跳转
*/
type IF_ACMPEQ struct{ base.BranchInstruction } // ==
type IF_ACMPNE struct{ base.BranchInstruction } // !=

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 != ref2 {
		base.Branch(frame, self.Offset)
	}
}
