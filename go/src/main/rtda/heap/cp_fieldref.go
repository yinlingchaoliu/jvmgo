package heap

import "main/classfile"

type FieldRef struct {
	MemberRef
	field *Field // 缓存解析后的字段指针
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

func (self *FieldRef) resolveFieldRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	field := lookupField(c, self.name, self.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}

	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.field = field
}

func lookupField(c *Class, name string, descriptor string) *Field {
	// 从当前类找
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	// 从实现的接口中找
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	// 从父类中找
	if field := lookupField(c.superClass, name, descriptor); field != nil {
		return field
	}
	return nil
}
