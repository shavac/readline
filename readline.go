// Wrapper around the GNU readline(3) library

package readline

// #include <stdio.h>
// #include <stdlib.h>
// #include <readline/readline.h>
// #include <readline/history.h>
import "C"
import "unsafe"

func ReadLine(prompt *string) *string {
	var p *C.char;

	//readline allows an empty prompt(NULL)
	if prompt != nil { p = C.CString(*prompt) }

	ret := C.readline(p);

	if p != nil { C.free(unsafe.Pointer(p)) }

	if ret == nil { return nil } //EOF

	s := C.GoString(ret);
	C.free(unsafe.Pointer(ret));
	return &s
}

func AddHistory(s string) {
	p := C.CString(s);
	C.add_history(p);
	C.free(unsafe.Pointer(p))
}
