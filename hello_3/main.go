package main

/*
#cgo CXXFLAGS: -std=c++14
#include "wrapper.h"
*/
import "C"
import "fmt"
//import "unsafe"

func main() {
    goMessage := "こんにちは！GoからC++への伝言です。"
	// 1-1. Goの文字列を Cの文字列（*C.char）に変換
	// ⚠️ このときCのメモリ（ヒープ）が新しく確保されます
	cMessage := C.CString(goMessage)
	// 1-2. 使い終わったら必ずメモリを解放する（お作法）
	// Goのガベージコレクション（GC）はCのメモリを掃除してくれません
	defer C.free_cpp_message(cMessage)
	// 1-3. C++の関数を呼び出す
	C.call_cpp_print(cMessage)

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