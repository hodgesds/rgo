package main

import "fmt"

// #cgo CFLAGS: -I/usr/share/R/include -I/usr/share/R
// #cgo LDFLAGS: -L/usr/lib/R -lR
// #include <stdlib.h>
// #include <Rembedded.h>
// #include <Rinterface.h>
import "C"

func main(){
    // for mor info see:
    // https://cran.r-project.org/doc/manuals/r-release/R-exts.html#Embedding-R-under-Unix_002dalikes
    fmt.Println("inside Go")

    argc := C.int(1);
    argv := make([]*C.char, 2)
    argv[0] = C.CString("rgo")
    argv[0] = C.CString("--silent")

    C.Rf_initEmbeddedR( argc, &argv[0] )

    //C.Rf_mainloop()
    C.R_ReplDLLinit()
    for {
        if (C.R_ReplDLLdo1()<=0){
            break
        }
    }

    // cleanup
    C.Rf_endEmbeddedR(0)
}
