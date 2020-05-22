package heap

import "unicode/utf16"

//string - ref
var internedStrings = map[string]*Object{}

//todo 支持javastring
// go string -> java.lang.String
func JString(loader *ClassLoader, goStr string) *Object {
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	//go 字符串 utf-8 => java utf-16
	chars := stringToUtf16(goStr)

	//加载 java类 char数组
	jChars := &Object{loader.LoadClass("[C"), chars}

	//加载java类 String
	jStr := loader.LoadClass("java/lang/String").NewObject()

	//反射设置
	jStr.SetRefVar("value", "[C", jChars)

	internedStrings[goStr] = jStr
	return jStr
}

// java.lang.String -> go string
func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	// utf16 -> utf-8
	return utf16ToString(charArr.Chars())
}

// utf8 -> utf16
func stringToUtf16(s string) []uint16 {
	runes := []rune(s)         // utf32
	return utf16.Encode(runes) // func Encode(s []rune) []uint16
}

// utf16 -> utf8
func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s) // func Decode(s []uint16) []rune
	return string(runes)
}

// todo
func InternString(jStr *Object) *Object {
	goStr := GoString(jStr)
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	internedStrings[goStr] = jStr
	return jStr
}
