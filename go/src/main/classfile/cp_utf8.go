package classfile

// 虚拟机规范
// CONSTANT_Utf8_info {
//    u1 tag;
//    u2 length;
//    u1 bytes[length];
// }
type ConstantUtf8Info struct {
	str string // 存放 MUTF-8 编码后的字符串
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := reader.readUint16()                            // u2 length
	self.str = decodeMUTF8(reader.readBytes(uint32(length))) // u1 bytes[length]
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes) // 为了简化，这里直接使用了UTF8编码（go默认编码），没有使用MUTF8
}
