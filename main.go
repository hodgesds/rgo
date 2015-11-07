package main

import (
	"fmt"
	"sync"
	//"os"
	//"unsafe"
)

// #cgo CFLAGS: -I/usr/share/R/include -I/usr/share/R
// #cgo LDFLAGS: -L/usr/lib/R -lR
// #include <stdlib.h>
// #include <Rembedded.h>
// #include <Rinterface.h>
import "C"

type RSession struct {
	Stdin chan interface{}
	Stout chan interface{}
}

func NewRSession() *RSession {
	in := make(chan interface{})
	out := make(chan interface{})
	rs := &RSession{in, out}
	go rs.start()
	return rs
}

func (rs *RSession) R_ShowMessage() {
}

//export R_Busy
func R_Busy(which C.int) {
	print("Busy?")
	//print(int(which))
}

//export R_ShowMessage
func R_ShowMessage(msg C.char) {
	print("go callback")
	print(msg)
}

//export R_ReadConsole
func R_ReadConsole(prompt *C.char, buf C.char, buflen C.int, hist C.int) C.int {
	print("READ ME!")
	prompt = C.CString("FOO> ")
	return C.int(0)
}

//export R_WriteConsole
func R_WriteConsole(buf C.char, buflen C.int) {
	print("WRITE ME!")
	print(C.GoString(&buf))
}

//export R_WriteConsoleEx
func R_WriteConsoleEx(buf C.char, buflen, otype C.int) {
	print("WRITE ME!")
}

//export R_ResetConsole
func R_ResetConsole() {
	print("RESET!!!")
}

//export R_FlushConsole
func R_FlushConsole() {
	print("FLUSH!")
}

//export R_ClearErrConsole
func R_ClearErrConsole() {
	print("CLEAR!")
}

//export R_ShowFiles
func R_ShowFiles(nfile C.int, file unsafe.Pointer, headers unsafe.Pointer, wtitle C.CString, del C.Rboolean, pager C.CString) {
}

//export R_ChooseFile
func R_ChooseFile(newFile C.int, buf C.CString, buflen C.int) C.int {
	return C.int(0)
}

//export R_EditFiles
func R_EditFiles(buf C.CString) C.int {
	return C.int(0)
}

//export R_loadhistory
func R_loadhistory(call, op, args, env C.SEXP) C.SEXP {
	print("loadhistory is not implemented")
	return C.R_NilValue
}

//export R_savehistory
func R_savehistory(call, op, args, env C.SEXP) C.SEXP {
	print("loadhistory is not implemented")
	return C.R_NilValue
}

//export R_addhistory
func R_addhistory(call, op, args, env C.SEXP) C.SEXP {
	print("loadhistory is not implemented")
	return C.R_NilValue
}

/*
//rexport R_Suicide
func R_Suicide(msg *unsafe.Pointer) {
}
*/

func (rs *RSession) start() {
	argc := C.int(1)
	argv := make([]*C.char, 7)
	argv[0] = C.CString("rgo")
	argv[1] = C.CString("--gui=none")
	argv[2] = C.CString("--no-save")
	argv[3] = C.CString("--no-readline")
	argv[4] = C.CString("--vanilla")
	argv[5] = C.CString("--slave")
	argv[6] = C.CString("--silent")

	//_ := os.Getenv("R_HOME") or C.get_R_HOME()
	C.Rf_initEmbeddedR(argc, &argv[0])

	C.R_ReplDLLinit()
	//C.Rf_mainloop()
	for {
		if C.R_ReplDLLdo1() <= 0 {
			break
		}
	}

}

// for reference:
// https://github.com/eddelbuettel/rinside/blob/master/src/RInside.cpp

//# Go string to C string; result must be freed with C.free
//# func C.CString(string) *C.char
//# C string to Go string
//# func C.GoString(*C.char) string
//# C string, length to Go string
//# func C.GoStringN(*C.char, C.int) string
//#	C pointer, length to Go []byte
//#	func C.GoBytes(unsafe.Pointer, C.int) []byte

func main() {
	fmt.Println("Yes we are inside of Go")
	var wg sync.WaitGroup
	wg.Add(1)
	i := 0
	for i < 1 {
		NewRSession()
		i++
	}
	wg.Wait()
}
