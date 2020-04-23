package classfile

// 标识类、接口、字段或方法不再推荐使用
type DeprecatedAttribute struct {
	MarkerAttribute
}

// 用于标识原文件中不存在的、由编译器生成的类成员
// 主要用于内部类或内部接口
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct{}

func (MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
