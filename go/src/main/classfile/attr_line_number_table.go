package classfile

// 作用：存放方法的行号信息
// 虚拟机规范:
/*
LineNumberTable_attribute {
    u2 attribute_name_index; 属性名索引，常量池索引，指向一个CONSTANT_Utf8_info常量
    u4 attribute_length; 属性长度
    u2 line_number_table_length; 方法行号表大小
    {
		u2 start_pc;
		u2 line_number;
    } line_number_table[line_number_table_length] 方法行号表
}
*/
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

//todo 代码行号
func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTable := make([]*LineNumberTableEntry, reader.readUint16())
	for i := range lineNumberTable {
		lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}

func (self *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(self.lineNumberTable) - 1; i >= 0; i-- {
		entry := self.lineNumberTable[i]
		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}
