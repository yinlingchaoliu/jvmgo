package heap

import "main/classfile"

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte // 方法字节码表
	argSlotCount uint // 参数个数
}

// 根据 classFile 创建 方法表
// classfile==cf  class==c
// 将 cf.method accessFlags、nameIndex、descriptorIndex 转化为具体的 accessFlags、name字符串、descriptor字符串 写入 c.method，将 cf.method 的 CodeAttribute 中的信息maxStack、maxLocals、code写入 c.method
func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calArgSlotCount() //参数
	}
	return methods
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}
}

func (self *Method) calArgSlotCount() {
	//解析函数描述
	parsedDescriptor := parseMethodDescriptor(self.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		self.argSlotCount++ // 编译器对于实例方法默认会在最前边添加一个 this 参数
	}
}

// getters
func (self *Method) MaxStack() uint {
	return self.maxStack
}
func (self *Method) MaxLocals() uint {
	return self.maxLocals
}
func (self *Method) Code() []byte {
	return self.code
}
func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}
func (self *Method) IsNative() bool {
	return 0 != self.accessFlags&ACC_NATIVE
}

func (self *ClassMember) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}

