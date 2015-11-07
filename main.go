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
// #include <R_ext/RStartup.h>
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

func (rs *RSession) R_Busy() {
}

func (rs *RSession) R_ReadConsole() {
}

func (rs *RSession) R_WriteConsole() {
}

func (rs *RSession) R_WriteConsoleEx() {
}

func (rs *RSession) R_ResetConsole() {
}

func (rs *RSession) R_FlushConsole() {
}

func (rs *RSession) R_ClearErrConsole() {
}

func (rs *RSession) R_ShowFiles() {
}

func (rs *RSession) R_ChooseFile() {
}

func (rs *RSession) R_EditFile() {
}

func (rs *RSession) R_EditFiles() {
}

func (rs *RSession) R_loadhistory() {
}

func (rs *RSession) R_savehistory() {
}

func (rs *RSession) R_addhistory() {
}

func (rs *RSession) R_Suicide() {
}

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
		//if C.R_ReplDLLdo1() <= 0 {
		//	break
		//}
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
