package constants

import (
	"main/instructions/base"
	"main/rtda"
)

type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// nothing to do
}
