package main

/*
#cgo CXXFLAGS: -std=c++14
#cgo LDFLAGS: -lsqlite3 
#include "wrapper.h"
*/
import "C"
import "fmt"
import "os"
import "strconv"

func main() {
    fmt.Println("#main-start")
    fmt.Println("全引数:", os.Args)
    if len(os.Args) < 2 {
        fmt.Println("error , argment none")
        return
    }
    var argment = os.Args[1]
    if argment == "add" {
        if len(os.Args) < 3 {
            fmt.Println("error , argment none")
            return
        }
        var title = os.Args[2]
        fmt.Println("title=", title)
        goMessage := title
        cMessage := C.CString(goMessage)
        // Goのガベージコレクション（GC）はCのメモリを掃除してくれません
        defer C.free_cpp_message(cMessage)
        C.todo_add(cMessage)   
    }
    if argment == "list" {
        C.todo_list()   
    }
    if argment == "del" {
        if len(os.Args) < 3 {
            fmt.Println("error , argment none")
            return
        }
        var id_str = os.Args[2]
        num, err := strconv.Atoi(id_str)
        if err != nil {
            fmt.Println("変換エラー:", err)
            return
        }
        fmt.Printf("型: %T, 値: %d\n", num, num) // 型: int, 値: 123        
        var id = num
        C.todo_delete(C.int(id))
    }




}