package classfile

// 作用：存放字节码等相关的方法信息
// 虚拟机规范:
/*
Code_attribute {
    u2 attribute_name_index; 属性名索引，常量池索引，指向一个CONSTANT_Utf8_info常量
    u4 attribute_length; 属性长度
    u2 max_stack; 操作数栈的最大深度
    u2 max_locals; 局部变量表的大小
    u4 code_length; 字节码表大小
    u1 code[code_length]; 字节码表
    u2 exception_table_length; 异常处理表大小
    {
		u2 start_pc;
		u2 end_pc;
		u2 hander_pc;
		u2 catch_type;
    } exception_table[exception_table_length] 异常处理表
    u2 attributes_count 属性表大小
    attribute_info attributes[attributes_count] 属性表
}
*/
type CodeAttribute struct {
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo

	cp ConstantPool
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	self.code = reader.readBytes(reader.readUint32())
	self.exceptionTable = readExceptionTable(reader);
	self.attributes = readAttributes(reader, self.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTable := make([]*ExceptionTableEntry, reader.readUint16())
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

// Getter
func (self *CodeAttribute) MaxStack() uint {
	return uint(self.maxStack)
}

func (self *CodeAttribute) MaxLocals() uint {
	return uint(self.maxLocals)
}

func (self *CodeAttribute) Code() []byte {
	return self.code
}
