package heap

import "main/classfile"

// 类符号引用
type ClassRef struct {
	SymRef
}

// 将 classfile.ConstantClassInfo 转化为 ClassRef
func newClassRef(cp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classInfo.Name()
	return ref
}
