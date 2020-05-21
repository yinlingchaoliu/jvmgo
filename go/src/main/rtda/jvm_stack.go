package rtda

//jvm栈
type Stack struct {
	maxSize uint   // 最多可容纳多少帧
	size    uint   // 当前栈中已经存放的栈帧数量
	_top    *Frame // 栈顶栈帧的指针
}

//新建jvm栈
func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

//栈帧入栈
func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil { //当栈顶不为空的话
		frame.lower = self._top //组建栈链表
	}
	self._top = frame // 将当前栈压入栈顶
	self.size++
}

//获得栈顶
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty")
	}
	//todo self._top
	return self._top
}

//栈帧出栈
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty")
	}
	top := self._top
	self._top = top.lower
	self.size--
	return top
}

//判断是否为空
func (self *Stack) IsEmpty() bool {
	return self._top == nil
}
