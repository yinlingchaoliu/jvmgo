package heap

//层次结构

// other是self的子类
// other实现了self接口
func (self *Class) isAssignableFrom(other *Class) bool {
	if self == other {
		return true
	}
	if self.IsInterface() {
		return other.isImplements(self)
	} else {
		return other.isSubClassOf(self)
	}
}

func (self *Class) isImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// self extends c
func (self *Class) isSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}