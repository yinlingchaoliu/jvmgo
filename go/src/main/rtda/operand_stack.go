package rtda

import (
	"main/rtda/heap"
	"math"
)

type OperandStack struct {
	size  uint   // 操作数栈中的 slot 个数 - 1, 即栈顶元素的位置
	slots []Slot // 操作数栈
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

/************************************** 基本类型 ***************************************/
// int
func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}

func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

// float
func (self *OperandStack) PushFloat(val float32) {
	self.slots[self.size].num = int32(math.Float32bits(val))
	self.size++
}

func (self *OperandStack) PopFloat() float32 {
	self.size--
	return math.Float32frombits(uint32(self.slots[self.size].num))
}

// long
func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].num = int32(val)         // 低32位
	self.slots[self.size+1].num = int32(val >> 32) // 高32位
	self.size += 2
}

func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size+1].num)
	return int64(high)<<32 | int64(low)
}

// double
func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}

func (self *OperandStack) PopDouble() float64 {
	return math.Float64frombits(uint64(self.PopLong()))
}

// byte char boolean short 省略，都可以转换成 int 处理

/************************************** 引用类型 ***************************************/
// 引用类型
func (self *OperandStack) PushRef(ref *heap.Object) {
	self.slots[self.size].ref = ref
	self.size++
}

func (self *OperandStack) PopRef() *heap.Object {
	self.size--
	return self.slots[self.size].ref
}

/************************************** slot ***************************************/
func (self *OperandStack) PushSlot(slot Slot) {
	self.slots[self.size] = slot
	self.size++
}

func (self *OperandStack) PopSlot() Slot {
	self.size--
	return self.slots[self.size]
}

func (self *OperandStack) GetRefFromTop(n uint) *heap.Object {
	return self.slots[self.size-1-n].ref
}

func (self *OperandStack) PushBoolean(val bool) {
	if val {
		self.PushInt(1)
	} else {
		self.PushInt(0)
	}
}

func (self *OperandStack) PopBoolean() bool {
	return self.PopInt() == 1
}

func NewOperandStack(maxStack uint) *OperandStack {
	return newOperandStack(maxStack)
}