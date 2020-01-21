package niall

// #include "niall.h"
import "C"

import (
	"bytes"
	"fmt"
)

var (
	WarningFunction = fmt.Println
	ErrorFunction   = fmt.Println
	PrintFunction   = fmt.Println
)

//export niall_go_warning
func niall_go_warning(msg *C.char) {
	WarningFunction("WARNING: " + C.GoString(msg))
}

//export niall_go_error
func niall_go_error(msg *C.char) {
	ErrorFunction("ERROR: " + C.GoString(msg))
}

//export niall_go_print
func niall_go_print(msg *C.char) {
	PrintFunction(C.GoString(msg))
}

// Make Niall learn from a string
func Learn(s string) {
	C.Niall_Learn(C.CString(s))
}

// Make Niall say something
func Talk() string {
	//size := 1024
	buf := C.malloc(C.sizeof_char * 1024)
	//buf, size := newCStringBuffer()
	C.Niall_Reply((*C.char)(buf), 1024)
	b := C.GoBytes(buf, 1024)
	if bytes.Contains(b, []byte(".")) {
		pos := bytes.Index(b, []byte("."))
		return string(b[:pos+1])
	}
	return string(b)
}

// This seeds the random number generator and clears the dictionary
func Init() {
	C.Niall_Init()
}

// This frees the allocated memory
func Quit() {
	C.Niall_Free()
}

// Clear the dictionary
func NewDictionary() {
	C.Niall_NewDictionary()
}

// List the contents of the dictionary
func ListDictionary() {
	C.Niall_ListDictionary()
}

// Load the dictionary from file
func LoadDictionary(filename string) {
	C.Niall_LoadDictionary(C.CString(filename))
}

// Save the dictionary to file
func SaveDictionary(filename string) {
	C.Niall_SaveDictionary(C.CString(filename))
}

// Correct the spelling
func CorrectSpelling(original, correct string) {
	C.Niall_CorrectSpelling(C.CString(original), C.CString(correct))
}
