package base

import (
	"main/rtda"
)

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextpc := pc + offset
	frame.SetNextPC(nextpc)
}
