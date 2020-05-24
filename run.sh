#!/bin/sh
set -ex

cd ./go
export GOPATH=$PWD

# 设置javahome
# export JAVA_HOME=/Library/Java/JavaVirtualMachines/jdk1.8.0_144.jdk/Contents/Home/

# 测试cmd
#go run main -version
#go run main -test "cmd" 12 344 567

# 测试搜索类路径
#go run main -test "classpath" java.lang.Object

# 测试加载classfile
#go run main -test "classfile" java.lang.Object

#测试 运行时数据
#go run main -test "rtda" "anystring"

#测试解释器
#go run main -test "interpret" -cp test/lib/example.jar jvmgo.book.ch05.GaussTest

##测试classloader
#go run main -test "classloader" -cp test/lib/example.jar jvmgo.book.ch06.MyObject

#测试函数调用返回
#go run main -verbose:class -verbose:inst  -test "return"  -cp test/lib/example.jar   jvmgo.book.ch07.InvokeDemo
#go run main -verbose:class -verbose:inst  -test "return"  -cp test/lib/example.jar   jvmgo.book.ch07.FibonacciTest

#数组功能测试  打印加载类 指令信息
#go run main   -test "array"  -cp test/lib/example.jar   jvmgo.book.ch08.BubbleSortTest

#测试字符串数组
#go run main   -test "string"  -cp test/lib/example.jar   jvmgo.book.ch01.HelloWorld

#go run main   -test "string"  -cp test/lib/example.jar   jvmgo.book.ch08.PrintArgs  'go jvm args' 'PrintArgs' 'Hello , World'

#测试本地方法调用
#go run main    -cp test/lib/example.jar   jvmgo.book.ch09.GetClassTest
#go run main    -cp test/lib/example.jar   jvmgo.book.ch09.StringTest
#go run main    -cp test/lib/example.jar   jvmgo.book.ch09.ObjectTest
#go run main    -cp test/lib/example.jar   jvmgo.book.ch09.CloneTest

#classloader 加载顺序
#go run main   -verbose:class  -cp test/lib/example.jar   jvmgo.book.ch09.TestLoadClass

#exception 异常处理
#go run main   -cp test/lib/example.jar   jvmgo.book.ch10.ParseIntTest  123
#go run main   -cp test/lib/example.jar   jvmgo.book.ch10.ParseIntTest  abc
go run main   -cp test/lib/example.jar   jvmgo.book.ch10.ParseIntTest


# todo 自动装箱和拆箱 第九章不处理
#go run main    -cp test/lib/example.jar   jvmgo.book.ch09.BoxTest

#go run jvmgo/ch02 java.lang.Object | grep -q "class data"
#go run jvmgo/ch03 java.lang.Object | grep -q "this class: java/lang/Object"
#go run jvmgo/ch04 java.lang.Object 2>&1 | grep -q "100"
#go run main -cp ../java/example.jar jvmgo.book.ch05.GaussTest 2>&1 | grep -q "5050"
#go run jvmgo/ch06 -cp ../java/example.jar jvmgo.book.ch06.MyObject | grep -q "32768"
#go run jvmgo/ch07 -cp ../java/example.jar jvmgo.book.ch07.FibonacciTest | grep -q "832040"
#go run jvmgo/ch08 -cp ../java/example.jar jvmgo.book.ch01.HelloWorld  | grep -q "Hello, world!"
#go run jvmgo/ch08 -cp ../java/example.jar jvmgo.book.ch08.PrintArgs foo bar | tr -d '\n' | grep -q "foobar"
#go run jvmgo/ch09 -cp ../java/example.jar jvmgo.book.ch09.GetClassTest | grep -q "Ljava.lang.String;"
#go run jvmgo/ch09 -cp ../java/example.jar jvmgo.book.ch09.StringTest | tr -d '\n' | grep -q "truefalsetrue"
#go run jvmgo/ch09 -cp ../java/example.jar jvmgo.book.ch09.ObjectTest | tr -d '\n' | grep -q "falsetrue"
#go run jvmgo/ch09 -cp ../java/example.jar jvmgo.book.ch09.CloneTest | grep -q "3.14"
#go run jvmgo/ch09 -cp ../java/example.jar jvmgo.book.ch09.BoxTest | grep -q "1, 2, 3"
#go run jvmgo/ch10 -cp ../java/example.jar jvmgo.book.ch10.ParseIntTest 123 | grep -q "123"
#go run jvmgo/ch10 -cp ../java/example.jar jvmgo.book.ch10.ParseIntTest abc 2>&1 | grep  'For input string: "abc"'
#go run jvmgo/ch10 -cp ../java/example.jar jvmgo.book.ch10.ParseIntTest 2>&1 | grep -q "at jvmgo"
#go run main -cp ../java/example.jar jvmgo.book.ch01.HelloWorld  | grep -q "Hello, world!"
#
echo OK