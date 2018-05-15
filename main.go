package main

import (
  "fmt"
  "runtime"
  "os"
  "./trans"
  )

/*
Go 代码中会使用到的 25 个关键字或保留字
break  default  func  interface  select
case  defer  go  map  struct
chan  else  goto  package  switch
const  fallthrough  if  range  type
continue  for  import  return  var

Go 语言还有 36 个预定义标识符
append  bool  byte  cap  close  complex  complex64  complex128  uint16
copy  false  float32  float64  imag  int  int8  int16  uint32
int32  int64  iota  len  make  new  nil  panic  uint64
print  println  real  recover  string  true  uint  uint8  uintptr
*/

/*
1. 在完成包的 import 之后，开始对常量、变量和类型的定义或声明。
2. 如果存在 init 函数的话，则对该函数进行定义（这是一个特殊的函数，每个含有该函数的包都会首先执行这个函数）。
3. 如果当前包是 main 包，则定义 main 函数。
4. 然后定义其余的函数，首先是类型的方法，接着是按照 main 函数中先后调用的顺序来定义相关函数，如果有很多函数，则可以按照字母顺序来进行排序。
*/

// 常量
const Name string = "golang"
const (
  Sunday = iota
  Monday
  Tuesday
  Wednesday
  Thursday
  Friday
  Saturday
)
// 全局变量,这种写法主要用于声明包级别的全局变量
var a string = "go语言"
var twoPi  = 2 * trans.Pi
// 一般类型声明
type goInt int
// 结构声明
type Learn struct {
}
// 接口声明
type Ilearn interface {
}
// main函数
func main() {
  fmt.Println("---learnGo func:")
  learnGo()
  fmt.Println("---goos func:")
  goos()
  fmt.Println("--- const Name:")
  fmt.Println(Name)
  fmt.Println("--- const Saturday:")
  fmt.Println(Saturday)
  var b string = "world"
  fmt.Println("hello",b)
  fmt.Println("--- init Pi:")
  fmt.Printf("2*Pi = %g\n", twoPi)
}
// 函数声明
func learnGo()  {
  // 当你在函数体内声明局部变量时，应使用简短声明语法 :=
  a := "learn go"
  fmt.Println(a)
}

func goos()  {
  var goos string = runtime.GOOS
  fmt.Printf("The operating system is: %s\n", goos)
  path := os.Getenv("PATH")
  fmt.Printf("Path is: %s\n", path)
}

/*
Go 程序的执行（程序启动）顺序如下：
  1. 按顺序导入所有被 main 包引用的其它包，然后在每个包中执行如下流程：
  2. 如果该包又导入了其它的包，则从第一步开始递归执行，但是每个包只会被导入一次。
  3. 然后以相反的顺序在每个包中初始化常量和变量，如果该包含有 init 函数的话，则调用该函数。
  4. 在完成这一切之后，main 也执行同样的过程，最后调用 main 函数开始执行程序。
*/
