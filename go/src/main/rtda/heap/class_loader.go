package heap

import (
	"fmt"
	"main/classfile"
	"main/classpath"
)

type ClassLoader struct {
	cp       *classpath.Classpath // 用于搜索和读取 class 文件
	classMap map[string]*Class    // 已经加载的类数据，key=全限定类名
}

// 创建一个类加载器
func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

// 把类数据加载到方法区
func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class // 类已经加载
	}
	return self.loadNonArrayClass(name) // 普通类的数据来自于class文件，数组类的数据是jvm在运行期间动态生成的
}

// 类加载过程
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name) // 1. 找到class文件并读取为 byte[]
	class := self.defineClass(data)     // 2. byte[] -> ClassFile -> Class，并放入方法区
	link(class)                         // 3. 进行链接
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)          // 递归加载父类
	resolveInterfaces(class)          // 递归加载接口类
	self.classMap[class.name] = class // 放入已加载列表
	return class
}

// byte[] -> ClassFile -> Class
func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName) // 递归加载父类
	}
}

func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)  // 验证
	prepare(class) // 准备
}

func verify(class *Class) {
	// todo
}

// 准备阶段：给类变量分配空间并给予初始值
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class) // 计算实例变量的个数，并分别标号
	calcStaticFieldSlotIds(class)   // 计算类变量的个数，并分别标号
	allocAndInitStaticVars(class)   // 为类变量分配空间并初始化
}

func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)

	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount // 子类的实例变量的编号要排在父类之后
	}

	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId ++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId // 实例变量所占空间个数
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)

	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId ++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId // 类变量所占空间个数
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount) // 为静态变量分配空间
	for _, field := range class.fields {
		// private static int y; // 普通静态变量，go语言本身自带默认值，所以不用做额外处理
		// private static final int x=1; // static+final，规定必须赋值，如果此处不直接赋值，需要在static块进行赋值
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

// private static final int x=1;
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.StaticVars() // 获取static变量
	cp := class.constantPool   // 获取常量池
	cpIndex := field.ConstantValueIndex()
	slotId := field.slotId

	if cpIndex > 0 { // private static final int x=1;
		switch field.descriptor { // https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html
		case "Z", "B", "C", "S", "I": // boolean byte char short integer
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val) // 设置静态变量初始值
		case "J": // long
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F": // float
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D": // double
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("todo") // todo

		}
	}
}
