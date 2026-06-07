package main

/*
#cgo CXXFLAGS: -std=c++14
#cgo LDFLAGS: -lcurl
#include "wrapper.h"
*/
import "C"
import "fmt"

func main() {
    fmt.Println("#main-start")
    C.test_curl();
}