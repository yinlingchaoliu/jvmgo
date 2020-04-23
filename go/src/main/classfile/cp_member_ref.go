package classfile

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}

// 虚拟机规范:
// ConstantFieldrefInfo 表示字段的符号引用
// ConstantMethodrefInfo 表示普通（非接口）方法的符号引用
// ConstantInterfaceMethodrefInfo 表示接口方法的符号引用
// ConstantFieldrefInfo {
//    u1 tag;
//    u2 class_index; 所在类，常量池索引，指向CONSTANT_Class_info常量
//    u2 name_and_type_index; 字段或方法名字和描述符，常量池索引，指向一个CONSTANT_Utf8_info常量
// }
// ConstantMethodrefInfo/ConstantInterfaceMethodrefInfo 与 ConstantFieldrefInfo 一样
type ConstantMemberrefInfo struct {
	classIndex       uint16 // 常量池索引，指向CONSTANT_Class_info常量
	nameAndTypeIndex uint16 // 常量池索引，指向CONSTANT_NameAndType_info常量

	cp ConstantPool
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}
