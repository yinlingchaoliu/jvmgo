#!/bin/sh
set -ex

cd ./go
export GOPATH=$PWD

# 设置javahome
# export JAVA_HOME=/Library/Java/JavaVirtualMachines/jdk1.8.0_144.jdk/Contents/Home/

## 测试cmd
#go run main -version
#go run main -test "cmd" 12 344 567
#
## 测试搜索类路径
#go run main -test "classpath" java.lang.Object
#
## 测试加载classfile
#go run main -test "classfile" java.lang.Object
#
##测试 运行时数据
#go run main -test "rtda" "anystring"
#
##测试解释器
#go run main -test "interpret" -cp test/lib/example.jar jvmgo.book.ch05.GaussTest
#
###测试classloader
#go run main -test "classloader" -cp test/lib/example.jar jvmgo.book.ch06.MyObject
#
##测试函数调用返回
#go run main -verbose:class -verbose:inst  -test "return"  -cp test/lib/example.jar   jvmgo.book.ch07.InvokeDemo
#go run main -verbose:class -verbose:inst  -test "return"  -cp test/lib/example.jar   jvmgo.book.ch07.FibonacciTest
#
##数组功能测试  打印加载类 指令信息
#go run main   -test "array"  -cp test/lib/example.jar   jvmgo.book.ch08.BubbleSortTest
#
##测试字符串数组
#go run main   -test "string"  -cp test/lib/example.jar   jvmgo.book.ch01.HelloWorld
#
#go run main   -test "string"  -cp test/lib/example.jar   jvmgo.book.ch08.PrintArgs  'go jvm args' 'PrintArgs' 'Hello , World'
#
##测试本地方法调用
#go run main  -test "string"  -cp test/lib/example.jar   jvmgo.book.ch09.GetClassTest
#go run main  -test "string"  -cp test/lib/example.jar   jvmgo.book.ch09.StringTest
#go run main  -test "string"  -cp test/lib/example.jar   jvmgo.book.ch09.ObjectTest
#go run main  -test "string"  -cp test/lib/example.jar   jvmgo.book.ch09.CloneTest
#
##classloader 加载顺序
#go run main  -test "string" -verbose:class  -cp test/lib/example.jar   jvmgo.book.ch09.TestLoadClass
#
##exception 异常处理
#go run main -test "string"  -cp test/lib/example.jar   jvmgo.book.ch10.ParseIntTest  123
#go run main -test "string"  -cp test/lib/example.jar   jvmgo.book.ch10.ParseIntTest  abc
#go run main -test "string"  -cp test/lib/example.jar   jvmgo.book.ch10.ParseIntTest

#system 打印hello world
go run main   -cp test/lib/example.jar   jvmgo.book.ch01.HelloWorld

echo OK