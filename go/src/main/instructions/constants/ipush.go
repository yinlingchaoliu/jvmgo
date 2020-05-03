package constants

import (
	"main/instructions/base"
	"main/rtda"
)

/*
BIPUSH: 从操作数中读取一个byte，并转换成int，推入操作数栈
SIPUSH: 从操作数中读取一个short，并转换成int，推入操作数栈
 */
type BIPUSH struct{ val int8 }

func (self *BIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(self.val))
}

type SIPUSH struct{ val int16 }

func (self *SIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(self.val))
}
