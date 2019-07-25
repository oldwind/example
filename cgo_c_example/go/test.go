package main

/*
#include <string.h>

*/
import "C"
import "fmt"

//export SayWhat
func SayWhat(str string) (*C.char, int32, int32) {
	cstr := C.CString(str)
	return cstr, int32(C.strlen(cstr)), 0
}

func main() {
	fmt.Println("it is a test")
}
