package classfile

import "encoding/binary"

//存在大小端问题
type ClassReader struct {
	data []byte
}

// u1 - 无符号1个字节（8位）
func (self *ClassReader) readUint8() uint8 {
	//取值
	val := self.data[0]
	//移动当前data位置
	self.data = self.data[1:]
	return val
}

// u2 - 无符号2个字节（16位）
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	//移动当前data位置
	self.data = self.data[2:]
	return val
}

// u4 - 无符号4个字节（32位）
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	//移动当前data位置
	self.data = self.data[4:]
	return val
}

// u8 - 无符号8个字节（64位）
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	//移动当前data位置
	self.data = self.data[8:]
	return val
}

// 读取interface表：
// 表头：eg. u2 interface_count
// 表项：eg. u2 interface[interface_count]
// 读取 exception_index_table 表
func (self *ClassReader) readUint16s() []uint16 {
	//第1位 数组大小 后续数组内容
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

// 读取指定字节个数据
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	//移动当前data位置
	self.data = self.data[n:]
	return bytes
}
