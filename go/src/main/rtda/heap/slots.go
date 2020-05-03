package heap

import "math"

// 变量（类变量、实例变量）
type Slot struct {
	num int32   // 基本类型（所有的基本类型都可以使用 int32 来处理）byte char short int float boolean -> 一个int32可表示；long double -> 需要两个int32表示
	ref *Object // 引用类型
}

type Slots []Slot

func newSlots(slotCount uint) Slots {
	if slotCount > 0 {
		return make([]Slot, slotCount)
	}
	return nil
}

/************************************** 基本类型 ***************************************/
// int
func (self Slots) SetInt(index uint, val int32) {
	self[index].num = val
}

func (self Slots) GetInt(index uint) int32 {
	return self[index].num
}

// float -> 转为 int
func (self Slots) SetFloat(index uint, val float32) {
	self[index].num = int32(math.Float32bits(val))
}

func (self Slots) GetFloat(index uint) float32 {
	return math.Float32frombits(uint32(self[index].num))
}

// long
func (self Slots) SetLong(index uint, val int64) {
	self[index].num = int32(val)         // 低32位
	self[index+1].num = int32(val >> 32) // 高32位
}

func (self Slots) GetLong(index uint) int64 {
	low := uint32(self[index].num)    // 低32位
	high := uint32(self[index+1].num) // 高32位
	return int64(high)<<32 | int64(low)
}

// double: 先转成long
func (self Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

func (self Slots) GetDouble(index uint) float64 {
	return math.Float64frombits(uint64(self.GetLong(index)))
}

/************************************** 引用类型 ***************************************/
// 引用类型
func (self Slots) SetRef(index uint, ref *Object) {
	self[index].ref = ref
}

func (self Slots) GetRef(index uint) *Object {
	return self[index].ref
}
