package main

/*
#cgo CXXFLAGS: -std=c++14
#include "wrapper.h"
*/
import "C"
import "fmt"

func main() {
	// C.関数名 で呼び出す。型はCの型になるので適宜キャスト
	result := C.call_cpp_add(C.int(10), C.int(20))
	
	fmt.Printf("C++からの計算結果: %d\n", int(result))
}