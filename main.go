package main

/*
#cgo LDFLAGS: -L./lib -lrustdemo
#include <stdlib.h>
#include "./lib/rustdemo.h"
*/
import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

func main() {

	argsWithoutProg := os.Args[1:]

	for _, arg := range argsWithoutProg {

		input := C.CString(arg)
		defer C.free(unsafe.Pointer(input))
		o := C.rustdemo(input)
		output := C.GoString(o)
		fmt.Printf("%s\n", output)

	}

}
