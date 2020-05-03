package heap

import (
	"fmt"
	"main/classfile"
)

// 常量项
type Constant interface{}

// 运行时常量池
type ConstantPool struct {
	class  *Class // 所属的类
	consts []Constant
}

// 把 classFile 中的常量池转化为运行时常量池
// ConstantUtf8Info 和 CONSTANT_NameAndType 会被解析到其他的常量里，
// eg. classfile.ConstantFieldrefInfo 中的 classIndex、nameAndTypeIndex -> FieldRef 中的 className、name、descriptor
func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, consts}

	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		// 字面量：整数、浮点数、字符串
		case *classfile.ConstantIntegerInfo:
			consts[i] = cpInfo.(*classfile.ConstantIntegerInfo).Value()
		case *classfile.ConstantFloatInfo:
			consts[i] = cpInfo.(*classfile.ConstantFloatInfo).Value()
		case *classfile.ConstantLongInfo:
			consts[i] = cpInfo.(*classfile.ConstantLongInfo).Value()
			i++
		case *classfile.ConstantDoubleInfo:
			consts[i] = cpInfo.(*classfile.ConstantDoubleInfo).Value()
			i++
		case *classfile.ConstantStringInfo:
			consts[i] = cpInfo.(*classfile.ConstantStringInfo).String()
			// 符号引用：类、字段、方法、接口方法
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
		case *classfile.ConstantFieldrefInfo:
			fieldrefInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
			consts[i] = newFieldRef(rtCp, fieldrefInfo)
		case *classfile.ConstantMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
			consts[i] = newMethodRef(rtCp, methodrefInfo)
		case *classfile.ConstantInterfaceMethodrefInfo:
			interfaceMethodrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, interfaceMethodrefInfo)
		}
	}
	return rtCp
}

// 根据索引返回常量项
func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No Constant at index %d", index))
}
