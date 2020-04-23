package classfile

// 作用：用于标识常量表达式的值
// 虚拟机规范:
// ConstantValue_attribute {
//    u2 attribute_name_index; 属性名索引，常量池索引，指向一个CONSTANT_Utf8_info常量
//    u4 attribute_length; 属性长度，恒等于2
//    u2 constantvalue_index; 属性值索引，常量池索引，指向一个CONSTANT_Integer_info常量
// }
type ConstantValueAttribute struct{
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}

