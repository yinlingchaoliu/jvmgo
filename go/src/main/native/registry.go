package native

import (
	"main/rtda"
)

/**
 * @author chentong
 * @desc 本地方法注册
 */

//本地方法定义为函数
type NativeMethod func(frame *rtda.Frame)

//定义函数数组
var registry = map[string]NativeMethod{}

//空方法
func emptyNativeMethod(frame *rtda.Frame) {
	// do nothing
}

//注册本地方法
func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

//寻找注册方法
func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}

	//如果是 object中 registerNatives 返回空方法，且此方法不会注册到registry中
	// todo initIDs
	if methodDescriptor == "()V" {
		if methodName == "registerNatives" || methodName == "initIDs" {
			return emptyNativeMethod
		}
	}

	return nil
}


