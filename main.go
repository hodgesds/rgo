package main

import (
	"fmt"
	//"sync"
	//"os"
	//"unsafe"
)

// #cgo CFLAGS: -I/usr/share/R/include -I/usr/share/R
// #cgo LDFLAGS: -L/usr/lib/R -lR
// #include <stdlib.h>
// #include <Rembedded.h>
// #include <Rinterface.h>
// #include <Rinternals.h>
// #include <Rversion.h>
// #include <R.h>
import "C"

//# for reference: pp 159ish
//# https://cran.r-project.org/doc/manuals/r-release/R-exts.pdf
//# interrupt handling: pp 145 R_ext/Utils.h

//export R_Busy
func R_Busy(which C.int) {
	if int(which) == 0 {
		print("thinking...\n")
	} else {
		print("that was easy!\n")
	}
}

/* refer to system.c */

//xxport R_ReadConsole
//func R_ReadConsole(prompt *C.char, buf *C.char, buflen C.int, hist C.int) C.int {
//	fmt.Println("gog:> ")
//	//return TrueReadConsole(prompt, (char *) buf, len, addtohistory)
//
//	var input string
//	fmt.Scanln(&input)
//	fmt.Println("in:", input)
//	//fmt.Scanln(buf)
//	buf = C.CString(input)
//
//	fmt.Println("BUF:", *buf)
//	fmt.Println("BUFLEN:", int(buflen))
//	buflen = C.int(len(C.GoString(buf)))
//	//buflen = C.int(len(buf))
//
//	//C.R_ProcessEvents()
//
//	return C.int(2)
//}

//export R_ProcessEvents
func R_ProcessEvents() {
	fmt.Println("event!")
}

/*
R_ProcessEvents();
    TrueWriteConsole(buf, len);
*/

//export R_WriteConsole
func R_WriteConsole(buf C.char, buflen C.int) {
	print("WRITE ME!\n")
	print(C.GoString(&buf))
}

//export R_WriteConsoleEx
func R_WriteConsoleEx(buf C.char, buflen C.int, otype C.int) {
	print("WRITE ME!\n")
	print(C.GoString(&buf))
}

//export R_ResetConsole
func R_ResetConsole() {
}

//export R_FlushConsole
func R_FlushConsole() {
	print("FLUSH!\n")
}

//export R_ClearErrConsole
func R_ClearErrConsole() {
	print("CLEAR!\n")
}

//xxport R_ShowFiles
//func R_ShowFiles(nfile C.int, file unsafe.Pointer, headers unsafe.Pointer, wtitle C.String, del C.Rboolean, pager C.String) {
//}

//xxport R_ChooseFile
//func R_ChooseFile(newFile C.int, buf C.String, buflen C.int) C.int {
//	return C.int(0)
//}

//xxport R_EditFiles
//func R_EditFiles(buf C.CString) C.int {
//	return C.int(0)
//}

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

func StartR() {
	argc := C.int(1)
	argv := make([]*C.char, 7)
	argv[0] = C.CString("rgo")
	argv[1] = C.CString("--gui=none")
	argv[2] = C.CString("--no-save")
	argv[3] = C.CString("--no-readline")
	argv[4] = C.CString("--vanilla")
	argv[5] = C.CString("--slave")
	argv[6] = C.CString("--silent")

	//print(C.ifp)
	//print(C.R_Consolefile)

	//C.R_SignalHandlers = C.int(0)

	C.Rf_initEmbeddedR(argc, &argv[0])

	C.R_ReplDLLinit()

	for {
		fmt.Println("loop")
		if C.R_ReplDLLdo1() <= 0 {
			C.Rf_endEmbeddedR(0)
			break
		} else {
			C.R_ProcessEvents()
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
	major := string(C.R_MAJOR)
	minor := string(C.R_MINOR)
	fmt.Println("About to start", major, minor)
	StartR()
}
