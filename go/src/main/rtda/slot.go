package rtda

import "main/rtda/heap"

type Slot struct {
	num int32        // 基本类型（所有的基本类型都可以使用 int32 来处理）
	ref *heap.Object // 引用类型
}
