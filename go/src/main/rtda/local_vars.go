package rtda

import (
"main/rtda/heap"
"math"
)

//todo 缺少Heap操作
type LocalVars []Slot // 本地变量表

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

/************************************** 基本类型 ***************************************/
// int
func (self LocalVars) SetInt(index uint, val int32) {
	self[index].num = val
}

func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

// float
func (self LocalVars) SetFloat(index uint, val float32) {
	self[index].num = int32(math.Float32bits(val))
}

func (self LocalVars) GetFloat(index uint) float32 {
	return math.Float32frombits(uint32(self[index].num))
}

// long
func (self LocalVars) SetLong(index uint, val int64) {
	self[index].num = int32(val)         // 低32位
	self[index+1].num = int32(val >> 32) // 高32位
}

func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)    // 低32位
	high := uint32(self[index+1].num) // 高32位
	return int64(high)<<32 | int64(low)
}

// double: 先转成long
func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

func (self LocalVars) GetDouble(index uint) float64 {
	return math.Float64frombits(uint64(self.GetLong(index)))
}

// byte char boolean short 省略，都可以转换成 int 处理

/************************************** 引用类型 ***************************************/
// 引用类型
func (self LocalVars) SetRef(index uint, ref *heap.Object) {
	self[index].ref = ref
}

func (self LocalVars) GetRef(index uint) *heap.Object {
	return self[index].ref
}

func (self LocalVars) SetSlot(index uint, slot Slot) {
	self[index] = slot
}

func (self LocalVars) GetThis() *heap.Object {
	return self.GetRef(0)
}

func (self LocalVars) GetBoolean(index uint) bool {
	return self.GetInt(index) == 1
}