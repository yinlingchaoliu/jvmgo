package classfile

// 作用：指出源文件名
// 虚拟机规范:
// SourceFile_attribute {
//    u2 attribute_name_index; 属性名索引，常量池索引，指向一个CONSTANT_Utf8_info常量
//    u4 attribute_length; 属性长度，恒等于2
//    u2 sourcefile_index; 属性值索引，常量池索引，指向一个CONSTANT_Utf8_info常量
// }
type SourceFileAttribute struct {
	sourceFileIndex uint16
	cp              ConstantPool
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}
