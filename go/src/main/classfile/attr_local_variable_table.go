package classfile

// 作用：存放方法的局部变量
// 虚拟机规范:
/*
LocalVariableTable_attribute {
    u2 attribute_name_index; 属性名索引，常量池索引，指向一个CONSTANT_Utf8_info常量
    u4 attribute_length; 属性长度
    u2 local_variable_table_length; 局部变量表大小
    {
		u2 start_pc;
		u2 length;
		u2 name_index; 局部变量名索引
		u2 descriptor_index; 局部变量名描述符索引
		u2 index; 局部变量在局部变量表中的位置
    } local_variable_table[local_variable_table_length] 局部变量表
}
*/
type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTable := make([]*LocalVariableTableEntry, reader.readUint16())
	for i := range localVariableTable {
		localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         reader.readUint16(),
			length:          reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index:           reader.readUint16(),
		}
	}
}
