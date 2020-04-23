package classfile

import (
	"math"
)

// 虚拟机规范
// CONSTANT_Integer_info {
//    u1 tag;
//    u4 bytes;
// }
// boolean/byte/short/char/int 都放在 ConstantIntegerInfo 中
type ConstantIntegerInfo struct {
	value int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	self.value = int32(reader.readUint32())
}

func (self *ConstantIntegerInfo) Value() int32 {
	return self.value
}
// 虚拟机规范
// CONSTANT_Float_info {
//    u1 tag;
//    u4 bytes;
// }
type ConstantFloatInfo struct {
	value float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	self.value = math.Float32frombits(reader.readUint32())
}

func (self *ConstantFloatInfo) Value() float32 {
	return self.value
}

// 虚拟机规范
// CONSTANT_Long_info {
//    u1 tag;
//    u4 high_bytes;
//    u4 low_bytes;
// }
type ConstantLongInfo struct {
	value int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	self.value = int64(reader.readUint64())
}

func (self *ConstantLongInfo) Value() int64 {
	return self.value
}
// 虚拟机规范
// CONSTANT_Double_info {
//    u1 tag;
//    u4 high_bytes;
//    u4 low_bytes;
// }
type ConstantDoubleInfo struct {
	value float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	self.value = math.Float64frombits(reader.readUint64())
}

func (self *ConstantDoubleInfo) Value() float64 {
	return self.value
}
