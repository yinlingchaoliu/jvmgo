package classfile

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type ClassFile struct {
	magic        uint32          // 魔数
	minorVersion uint16          // 次版本号
	majorVersion uint16          // 主版本号
	constantPool ConstantPool    // 常量池
	accessFlags  uint16          // 类访问标志
	thisClass    uint16          // 类常量池索引
	superClass   uint16          // 父类常量池索引（只有Object为0）
	interfaces   []uint16        // 接口常量池索引表
	fields       []*MemberInfo   // 字段表
	methods      []*MemberInfo   // 方法表
	attributes   []AttributeInfo // 属性表
}

// 将 []byte 转换成 ClassFile
func Parse(classData []byte) (cf *ClassFile, err error) {
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

// 使用 ClassReader 从 ClassReader 中读取内容，赋值给 ClassFile 的各个属性
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

// 读取并检查魔数
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32() // 读取魔数
	if magic != 0xCAFEBABE { // 检查魔数
		panic("java.lang.ClassFormatError: magic!")
	}
}

// 读取并检查主次版本号
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16() // 次版本号
	self.majorVersion = reader.readUint16() // 主版本号
	switch self.majorVersion {
	case 45: // jdk1.0 ~ jdk1.1，次版本号不为0
		return
	case 46, 47, 48, 49, 50, 51, 52, 53, 54: // jdk1.2 ~ jdk10，此版本号都为0
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// 从常量池查找类名
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

// 从常量池查找父类名
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" // java.lang.Object 没有父类
}

// 从常量池查找接口名
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

// Getter
func (self *ClassFile) Magic() uint32 {
	return self.magic
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) ThisClass() uint16 {
	return self.thisClass
}

func (self *ClassFile) SuperClass() uint16 {
	return self.superClass
}

func (self *ClassFile) Interfaces() []uint16 {
	return self.interfaces
}

func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}
