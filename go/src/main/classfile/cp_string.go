package classfile

// 虚拟机规范
// CONSTANT_String_info {
//    u1 tag;
//    u2 string_index; 不存放字符串本身，存常量池索引，指向一个CONSTANT_Utf8_info常量
// }
type ConstantStringInfo struct {
	stringIndex uint16 // 指向一个CONSTANT_Utf8_info的常量池索引
	cp          ConstantPool
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}