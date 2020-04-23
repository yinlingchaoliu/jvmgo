package classfile

// 虚拟机规范: CONSTANT_NameAndType_info 给出字段或方法的名称和描述符
// CONSTANT_Class_info 与 CONSTANT_NameAndType_info 可以唯一确定一个字段或方法
// CONSTANT_NameAndType_info {
//    u1 tag;
//    u2 name_index; 字段或方法名，常量池索引，指向一个CONSTANT_Utf8_info常量
//    u2 descriptor_index; 字段或方法描述符，常量池索引，指向一个CONSTANT_Utf8_info常量
// }
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16 // 字段或方法名 eg.main
	descriptorIndex uint16 // 字段或方法描述符 eg. java.lang.Object -> Ljava.lang.Object; void main(string[] args) -> ([Ljava.lang.String;)V
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
