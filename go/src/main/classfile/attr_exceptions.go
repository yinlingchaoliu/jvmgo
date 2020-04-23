package classfile

// 作用：记录方法抛出的异常表
// 虚拟机规范:
/*
Exceptions_attribute {
    u2 attribute_name_index; 属性名索引，常量池索引，指向一个CONSTANT_Utf8_info常量
    u4 attribute_length; 属性长度
    u2 number_of_exceptions; 异常表大小
    u2 exception_index_table[number_of_exceptions]; 异常表
}
*/
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionsAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}

func (self *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}