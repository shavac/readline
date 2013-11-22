// Wrapper around the GNU readline(3) library

package readline

/*
#cgo darwin CFLAGS: -I/opt/local/include
#cgo darwin LDFLAGS: -L/opt/local/lib
#cgo LDFLAGS: -lreadline

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "readline/readline.h"
#include "readline/history.h"

extern char** completerShim(char*, int, int);

*/
import "C"
import "unsafe"
import "syscall"
import "os"
import "os/signal"

//SIGWINCH handling is here.
func init() {
	C.rl_catch_sigwinch = 0
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGWINCH)
	go func() {
		for sig := range c {
			switch sig {
			case syscall.SIGWINCH:
				Resize()
			default:

			}
		}
	}()

	C.rl_attempted_completion_function = (*C.rl_completion_func_t)(C.completerShim)
}

func Resize() {
	C.rl_resize_terminal()
}

func ReadLine(prompt *string) *string {
	var p *C.char

	//readline allows an empty prompt(NULL)
	if prompt != nil {
		p = C.CString(*prompt)
	}

	ret := C.readline(p)

	if p != nil {
		C.free(unsafe.Pointer(p))
	}

	if ret == nil {
		return nil
	} //EOF

	s := C.GoString(ret)
	C.free(unsafe.Pointer(ret))
	return &s
}

func AddHistory(s string) {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	C.add_history(p)
}

// Parse and execute single line of a readline init file.
func ParseAndBind(s string) {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	C.rl_parse_and_bind(p)
}

// Parse a readline initialization file.
// The default filename is the last filename used.
func ReadInitFile(s string) error {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	errno := C.rl_read_init_file(p)
	if errno == 0 {
		return nil
	}
	return syscall.Errno(errno)
}

// Load a readline history file.
// The default filename is ~/.history.
func ReadHistoryFile(s string) error {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	errno := C.read_history(p)
	if errno == 0 {
		return nil
	}
	return syscall.Errno(errno)
}

var (
	HistoryLength = -1
)

// Save a readline history file.
// The default filename is ~/.history.
func WriteHistoryFile(s string) error {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	errno := C.write_history(p)
	if errno == 0 && HistoryLength >= 0 {
		errno = C.history_truncate_file(p, C.int(HistoryLength))
	}
	if errno == 0 {
		return nil
	}
	return syscall.Errno(errno)
}

// Set the readline word delimiters for tab-completion
func SetCompleterDelims(break_chars string) {
	p := C.CString(break_chars)
	//defer C.free(unsafe.Pointer(p))
	C.free(unsafe.Pointer(C.rl_completer_word_break_characters))
	C.rl_completer_word_break_characters = p
}

// Get the readline word delimiters for tab-completion
func GetCompleterDelims() string {
	cstr := C.rl_completer_word_break_characters
	delims := C.GoString(cstr)
	return delims
}

var completer func(string, int, int) []string

// SetCompletionFunction sets the function that will be used when the user
// invokes completion.
func SetCompletionFunction(c func(string, int, int)[]string) {
	completer = c
}

//export ProcessCompletion
func ProcessCompletion(textC *C.char, start, end int) **C.char {
	if completer == nil {
		return nil
	}

	text := C.GoString(textC)

	results := completer(text, start, end)

	if len(results) == 0 {
		return nil
	}

	var c *C.char
	ptrSize := unsafe.Sizeof(c)
	ptr := C.malloc(C.size_t(len(results) + 1) * C.size_t(ptrSize))

	for idx, value := range results {
		element := (**C.char)(unsafe.Pointer(uintptr(ptr) + uintptr(idx)*ptrSize))
		*element = (*C.char)(C.CString(value))
	}
	endElement := (**C.char)(unsafe.Pointer(uintptr(ptr) + uintptr(len(results))*ptrSize))
	*endElement = nil

	return (**C.char)(ptr)
}
