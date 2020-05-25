package misc

import (
	"main/native"
	"main/rtda"
	"main/rtda/heap"
)

const miscUnsafe = "sun/misc/Unsafe"

func init() {
	native.Register(miscUnsafe, "arrayBaseOffset", "(Ljava/lang/Class;)I", arrayBaseOffset)
	native.Register(miscUnsafe, "arrayIndexScale", "(Ljava/lang/Class;)I", arrayIndexScale)
	native.Register(miscUnsafe, "addressSize", "()I", addressSize)
	native.Register(miscUnsafe, "objectFieldOffset", "(Ljava/lang/reflect/Field;)J", objectFieldOffset)
	native.Register(miscUnsafe, "compareAndSwapObject", "(Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z", compareAndSwapObject)


}

// public native int arrayBaseOffset(Class<?> type);
// (Ljava/lang/Class;)I
func arrayBaseOffset(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PushInt(0) // todo
}

// public native int arrayIndexScale(Class<?> type);
// (Ljava/lang/Class;)I
func arrayIndexScale(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PushInt(1) // todo
}

// public native int addressSize();
// ()I
func addressSize(frame *rtda.Frame) {
	// vars := frame.LocalVars()
	// vars.GetRef(0) // this

	stack := frame.OperandStack()
	stack.PushInt(8) // todo unsafe.Sizeof(int)
}

func objectFieldOffset(frame *rtda.Frame) {
	vars := frame.LocalVars()
	jField := vars.GetRef(1)

	offset := jField.GetIntVar("slot", "I")

	stack := frame.OperandStack()
	stack.PushLong(int64(offset))
}

// public final native boolean compareAndSwapObject(Object o, long offset, Object expected, Object x)
// (Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z
func compareAndSwapObject(frame *rtda.Frame) {
	vars := frame.LocalVars()
	obj := vars.GetRef(1)
	fields := obj.Data()
	offset := vars.GetLong(2)
	expected := vars.GetRef(4)
	newVal := vars.GetRef(5)

	// todo
	if anys, ok := fields.(heap.Slots); ok {
		// object
		swapped := _casObj(obj, anys, offset, expected, newVal)
		frame.OperandStack().PushBoolean(swapped)
	} else if objs, ok := fields.([]*heap.Object); ok {
		// ref[]
		swapped := _casArr(objs, offset, expected, newVal)
		frame.OperandStack().PushBoolean(swapped)
	} else {
		// todo
		panic("todo: compareAndSwapObject!")
	}
}
func _casObj(obj *heap.Object, fields heap.Slots, offset int64, expected, newVal *heap.Object) bool {
	current := fields.GetRef(uint(offset))
	if current == expected {
		fields.SetRef(uint(offset), newVal)
		return true
	} else {
		return false
	}
}
func _casArr(objs []*heap.Object, offset int64, expected, newVal *heap.Object) bool {
	current := objs[offset]
	if current == expected {
		objs[offset] = newVal
		return true
	} else {
		return false
	}
}
