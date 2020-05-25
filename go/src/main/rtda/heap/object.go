package heap

// 表示实例
type Object struct {
	class  *Class
	//fields Slots
	data interface{} //todo 数组对象 interface{}  = void*
	extra interface{}  //todo native 记录额外信息 class
}

// 新创建的实例对象需要赋初值，go默认赋了
func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		data: newSlots(class.instanceSlotCount),
	}
}

// getters
func (self *Object) Class() *Class {
	return self.class
}

//todo 数组调整
func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(self.class)
}

//支持反射
// reflection
func (self *Object) GetRefVar(name, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}

func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (self *Object) SetIntVar(name, descriptor string, val int32) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetInt(field.slotId, val)
}

func (self *Object) GetIntVar(name, descriptor string) int32 {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetInt(field.slotId)
}

//获得class
func (self *Object) Extra() interface{} {
	return self.extra
}

//设置class todo exception
func (self *Object) SetExtra(extra interface{}) {
	self.extra = extra
}

func (self *Object) Data() interface{} {
	return self.data
}