package heap

// 符号引用基类
type SymRef struct {
	cp        *ConstantPool // 符号引用所在的常量池
	className string        // 类的全限定名
	class     *Class        // 符号引用所属的类
}

func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}
