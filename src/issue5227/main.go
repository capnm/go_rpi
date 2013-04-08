// https://code.google.com/p/go/issues/detail?id=5227
package main

//#include "init.h"
import "C"

func main() {
	C.init()
}

func selectfont() C.Fontinfo {
	return C.SansTypeface
}
