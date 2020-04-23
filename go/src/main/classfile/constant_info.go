package classfile

// tag 枚举。
// 每一个常量项都由 tag 和 常量项内容组成
const (
	CONSTANT_Utf8               = 1  // UTF-8 编码的字符串
	CONSTANT_Integer            = 3  // 整型字面量
	CONSTANT_Float              = 4  // 浮点型字面量
	CONSTANT_Long               = 5  // 长整型字面量
	CONSTANT_Double             = 6  // 双精度浮点型字面量
	CONSTANT_Class              = 7  // 类或接口的符号引用
	CONSTANT_String             = 8  // 字符串类型字面量
	CONSTANT_Fieldref           = 9  // 字段的符号引用
	CONSTANT_Methodref          = 10 // 类中方法的符号引用
	CONSTANT_InterfaceMethodref = 11 // 接口中方法的符号引用
	CONSTANT_NameAndType        = 12 // 字段或方法的部分符号引用（名字和描述符）
	CONSTANT_MethodHandle       = 15 // 方法句柄
	CONSTANT_MethodType         = 16 // 方法类型
	CONSTANT_InvokeDynamic      = 18 // 表示一个动态方法调用点
)

// 常量项接口
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

// 模板方法 - 模板模式
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()     // 读取tag
	c := newConstantInfo(tag, cp) // 创建常量项
	c.readInfo(reader)            // 初始化常量项
	return c
}

// 根据tag创建常量项
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
