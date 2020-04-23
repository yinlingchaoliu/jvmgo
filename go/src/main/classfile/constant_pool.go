package classfile

// 常量池：即常量项数组
type ConstantPool []ConstantInfo

// 读取常量项，构建并初始化常量池
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16()); // 读取常量池大小 n
	cp := make([]ConstantInfo, cpCount)  // 创建常量池
	for i := 1; i < cpCount; i++ { // 常量池索引从1开始，0要空出来（常量池大小 n-1）
		cp[i] = readConstantInfo(reader, cp) // 读取常量项
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ // long double 占两个索引位置
		}
	}
	return cp
}

// 根据索引从常量池获取常量项
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

// 根据索引从常量池查找字段或方法的名字和描述符
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

// 根据索引从常量池查找类名
func (self ConstantPool) getClassName(index uint16) string {
	cInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(cInfo.nameIndex)
}

// 从常量池查找utf8字符串
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
