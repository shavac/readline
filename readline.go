// Wrapper around the GNU readline(3) library

package readline

// TODO:
//  implement a go-oriented command completion

/*
 #cgo darwin CFLAGS: -I/opt/local/include
 #cgo darwin LDFLAGS: -L/opt/local/lib
 #cgo LDFLAGS: -lreadline

 #include <stdio.h>
 #include <stdlib.h>
 #include <string.h>
 #include "readline/readline.h"
 #include "readline/history.h"

 char* _go_readline_strarray_at(char **strarray, int idx) 
 {
   return strarray[idx];
 }

 int _go_readline_strarray_len(char **strarray)
 {
   int sz = 0;
   while (strarray[sz] != NULL) {
     sz += 1;
   }
   return sz;
 }
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
				;
			}
		}
	}()
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

//
func CompletionMatches(text string, cbk func(text string, state int) string) []string {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))
	c_cbk := (*C.rl_compentry_func_t)(unsafe.Pointer(&cbk))
	c_matches := C.rl_completion_matches(c_text, c_cbk)
	n_matches := int(C._go_readline_strarray_len(c_matches))
	matches := make([]string, n_matches)
	for i := 0; i < n_matches; i++ {
		matches[i] = C.GoString(C._go_readline_strarray_at(c_matches, C.int(i)))
	}
	return matches
}

//
func SetAttemptedCompletionFunction(cbk func(text string, start, end int) []string) {
	c_cbk := (*C.rl_completion_func_t)(unsafe.Pointer(&cbk))
	C.rl_attempted_completion_function = c_cbk
}

/* EOF */
