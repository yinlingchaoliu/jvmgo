package classfile

// 虚拟机规范: CONSTANT_Class_info 是类或者接口的符号引用
// CONSTANT_Class_info {
//    u1 tag;
//    u2 name_index; 常量池索引，指向一个CONSTANT_Utf8_info常量
// }
// 类索引、父类索引、接口表中的接口索引指向的都是CONSTANT_Class_info常量
type ConstantClassInfo struct {
	nameIndex uint16
	cp        ConstantPool
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
