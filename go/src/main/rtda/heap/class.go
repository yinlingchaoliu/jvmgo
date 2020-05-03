package heap

import (
	"main/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16        // 类访问标志
	name              string        // 类名（全限定）
	superClassName    string        // 父类名（全限定），eg. java/lang/Object
	interfaceNames    []string      // 接口名（全限定）
	constantPool      *ConstantPool // 运行时常量池
	fields            []*Field      // 字段表
	methods           []*Method     // 方法表
	loader            *ClassLoader  // 类加载器
	superClass        *Class        // 父类指针
	interfaces        []*Class      // 实现的接口指针
	instanceSlotCount uint          // 存放实例变量占据的空间大小（包含从父类继承来的实例变量）（其中long和double占两个slot）
	staticSlotCount   uint          // 存放类变量占据的空间大小（只包含当前类的类变量）（其中long和double占两个slot）
	staticVars        Slots         // 存放静态变量
	initStarted       bool
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() ||
		self.getPackageName() == other.getPackageName()
}

func (self *Class) getPackageName() string {
	// eg. java/lang/String
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}
func (self *Class) getStaticMethod(name string, descriptor string) *Method {
	for _, m := range self.methods {
		if m.IsStatic() && m.Name() == name && m.Descriptor() == descriptor {
			return m
		}
	}
	return nil
}

func (self *Class) GetClinitMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

// getters
func (self *Class) Name() string {
	return self.name
}
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}
func (self *Class) InitStarted() bool {
	return self.initStarted
}
func (self *Class) SuperClass() *Class {
	return self.superClass
}

// setter
func (self *Class) StartInit() {
	self.initStarted = true
}

