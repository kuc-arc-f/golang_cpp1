package main

/*
#cgo CXXFLAGS: -std=c++14
#include "wrapper.h"
*/
import "C"
import "fmt"

func main() {
	// 1. C++から文字列のポインタ（アドレス）を受け取る
	cPtr := C.get_cpp_message()

	// 2. 使い終わったら「必ず」C++側の解放関数を呼ぶように予約
	// これを忘れると、関数を呼ぶたびにメモリが漏れ続けます（メモリリーク）
	defer C.free_cpp_message(cPtr)

	// 3. Cの文字列を Goの string にコピーする
	// ⚠️ ここでGoが安全に管理できるメモリにデータが複製されます
	goStr := C.GoString(cPtr)

	// 4. あとはGoの文字列として自由に使える
	fmt.Println("受け取った文字列:", goStr)
}