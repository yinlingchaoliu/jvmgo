package heap

import "main/classfile"

// Field 与 Method 的父类，不是 Class 的父类
type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class // 所属的类
}

// 从 classFile 中复制数据
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

// d 是否可以访问 self(字段或方法)
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	// self 是 public
	if self.IsPublic() {
		return true
	}
	c := self.class
	// self 是 protected，则只有 d 是 self所在的class c的子类或者同一个包可以访问
	// 注意 protected 不只是子类级别，同包也可访问
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) || c.getPackageName() == d.getPackageName()
	}
	// self 是 default 级别
	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}
func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}

// 判断是否为static类
func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}

// getters
func (self *ClassMember) Name() string {
	return self.name
}
func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) Class() *Class {
	return self.class
}