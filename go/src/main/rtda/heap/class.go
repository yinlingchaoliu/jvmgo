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
	sourceFile        string		// 源文件名字
	loader            *ClassLoader  // 类加载器
	superClass        *Class        // 父类指针
	interfaces        []*Class      // 实现的接口指针
	instanceSlotCount uint          // 存放实例变量占据的空间大小（包含从父类继承来的实例变量）（其中long和double占两个slot）
	staticSlotCount   uint          // 存放类变量占据的空间大小（只包含当前类的类变量）（其中long和double占两个slot）
	staticVars        Slots         // 存放静态变量
	initStarted       bool			//类初始化标志 todo
	jClass            *Object		// java.lang.Class实例
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
	class.sourceFile = getSourceFile(cf)
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
	return self.GetStaticMethod("main", "([Ljava/lang/String;)V")
}
func (self *Class) getStaticMethod(name string, descriptor string) *Method {
	for _, m := range self.methods {
		if m.IsStatic() && m.Name() == name && m.Descriptor() == descriptor {
			return m
		}
	}
	return nil
}

//todo 初始化执行方法 <clinit>
func (self *Class) GetClinitMethod() *Method {
	return self.GetStaticMethod("<clinit>", "()V")
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

//todo 查到加载自身classloader
func (self *Class) Loader() *ClassLoader {
	return self.loader
}

// todo 增加arrayclass
func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}


func (self *Class) isJlObject() bool {
	return self.name == "java/lang/Object"
}

func (self *Class) isJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}
func (self *Class) isJioSerializable() bool {
	return self.name == "java/io/Serializable"
}

// todo 反射支持
// 根据字段名称和描述 查找
func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic &&
				field.name == name &&
				field.descriptor == descriptor {

				return field
			}
		}
	}
	return nil
}

//寻找函数
func (self *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for c := self; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {

				return method
			}
		}
	}
	return nil
}

// todo 支持反射 reflection
func (self *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := self.getField(fieldName, fieldDescriptor, true)
	return self.staticVars.GetRef(field.slotId)
}

func (self *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := self.getField(fieldName, fieldDescriptor, true)
	self.staticVars.SetRef(field.slotId, ref)
}

//获得object
func (self *Class) JClass() *Object {
	return self.jClass
}

//获得JavaName  java.lang.String
func (self *Class) JavaName() string {
	return strings.Replace(self.name, "/", ".", -1)
}

//是否为基本数据类型
func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]
	return ok
}

func (self *Class) GetStaticMethod(name, descriptor string) *Method {
	return self.getMethod(name, descriptor, true)
}

func (self *Class) GetInstanceMethod(name, descriptor string) *Method {
	return self.getMethod(name, descriptor, false)
}

func (self *Class) SourceFile() string {
	return self.sourceFile
}

//todo execption 获得源文件名称
func getSourceFile(cf *classfile.ClassFile) string {
	if sfAttr := cf.SourceFileAttribute(); sfAttr != nil {
		return sfAttr.FileName()
	}
	return "Unknown"
}