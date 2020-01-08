package main

/*
#cgo CFLAGS: -g -Wall
#include <stdlib.h>
#include "printstring.h"
*/
import "C"

import (
    "fmt"
    "unsafe"
)

func main() {

	p := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(p))

        size := C.printstring(C.CString("Hello World!"), (*C.char)(p))
        s := C.GoBytes(p, size)
        fmt.Println("in go: " + string(s))
}
