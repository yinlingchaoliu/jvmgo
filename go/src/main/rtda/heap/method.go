package heap

import (
	"main/classfile"
)

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte // 方法字节码表
	argSlotCount uint // 参数个数
	exceptionTable ExceptionTable //异常处理表
	lineNumberTable * classfile.LineNumberTableAttribute
}

// 根据 classFile 创建 方法表
// classfile==cf  class==c
// 将 cf.method accessFlags、nameIndex、descriptorIndex 转化为具体的 accessFlags、name字符串、descriptor字符串 写入 c.method，将 cf.method 的 CodeAttribute 中的信息maxStack、maxLocals、code写入 c.method

//新建一组方法
func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	//method.calArgSlotCount() //参数
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
		// todo 代码行号
		self.lineNumberTable = codeAttr.LineNumberTableAttribute()
		//todo exception 增加异常处理
		self.exceptionTable = newExceptionTable(codeAttr.ExceptionTable(),self.class.constantPool)
	}
}

func (self *Method) calcArgSlotCount(paramTypes []string) {
	for _, paramType := range paramTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		self.argSlotCount++ // `this` reference
	}
}

//todo 0xfe  invoke_native
func (self *Method) injectCodeAttribute(returnType string) {
	self.maxStack = 4 // todo
	self.maxLocals = self.argSlotCount
	switch returnType[0] {
	case 'V':
		self.code = []byte{0xfe, 0xb1} // return
	case 'L', '[':
		self.code = []byte{0xfe, 0xb0} // areturn
	case 'D':
		self.code = []byte{0xfe, 0xaf} // dreturn
	case 'F':
		self.code = []byte{0xfe, 0xae} // freturn
	case 'J':
		self.code = []byte{0xfe, 0xad} // lreturn
	default:
		self.code = []byte{0xfe, 0xac} // ireturn
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

//寻找异常处理句柄
func (self *Method) FindExceptionHandler(exClass *Class, pc int) int {
	handler := self.exceptionTable.findExceptionHandler(exClass, pc)
	if handler != nil {
		return handler.handlerPc
	}
	return -1
}

//获得行号
func (self *Method) GetLineNumber(pc int) int {
	if self.IsNative() {
		return -2
	}
	if self.lineNumberTable == nil {
		return -1
	}
	return self.lineNumberTable.GetLineNumber(pc)
}