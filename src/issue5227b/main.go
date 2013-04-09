//https://code.google.com/p/go/issues/detail?id=5227

package main

/*
typedef struct {
        int Count;
} Fontinfo;

Fontinfo SansTypeface;

extern void init();

Fontinfo loadfont() {
        Fontinfo f;
        return f;
}

void init() {
        SansTypeface = loadfont();
}
*/
import "C"

func main() {
	C.init()
}

func selectfont() C.Fontinfo {
	return C.SansTypeface
}
