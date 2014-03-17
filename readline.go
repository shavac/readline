// Copyright 2010-2014 go-readline authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Wrapper around the GNU readline(3) library

package readline

/*
#cgo darwin CFLAGS: -I/usr/local/Cellar/readline/6.2.4/include/
#cgo darwin LDFLAGS: -L/usr/local/Cellar/readline/6.2.4/lib/
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

// init handles window resizing on SIGWINCH.
func init() {
	C.rl_catch_sigwinch = 0
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGWINCH)
	go func() {
		for sig := range c {
			switch sig {
			case syscall.SIGWINCH:
				ResizeTerminal()
			default:

			}
		}
	}()
}

// ResizeTerminal updates the internal screen size by reading values from
// the kernel.
func ResizeTerminal() {
	C.rl_resize_terminal()
}

// Readline prints the prompt and reads a line from the standard input.
func Readline(prompt *string) *string {
	var p *C.char

	// readline allows an empty prompt (NULL)
	if prompt != nil {
		p = C.CString(*prompt)
	}

	ret := C.readline(p)

	if p != nil {
		C.free(unsafe.Pointer(p))
	}

	if ret == nil {
		return nil
	} // EOF

	s := C.GoString(ret)
	C.free(unsafe.Pointer(ret))
	return &s
}

// AddHistory adds a string to the end of the history list.
func AddHistory(s string) {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	C.add_history(p)
}

// ParseAndBind parses line as if it had been read from the inputrc file
// and performs any key bindings and variable assignments found.
func ParseAndBind(s string) {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	C.rl_parse_and_bind(p)
}

// ReadInitFile reads keybindings and variable assignments from filename.
// The default filename is the last filename used.
func ReadInitFile(filename string) error {
	p := C.CString(filename)
	defer C.free(unsafe.Pointer(p))
	if errno := C.rl_read_init_file(p); errno != 0 {
		return syscall.Errno(errno)
	}
	return nil
}

// ReadHistory loads a readline history file.
// The default filename is ~/.history.
func ReadHistoryFile(s string) error {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	if errno := C.read_history(p); errno != 0 {
		return syscall.Errno(errno)
	}
	return nil
}

var (
	HistoryLength = -1 // Maximum number of lines in the history file.
)

// WriteHistory saves a readline history file.
// The default filename is ~/.history.
func WriteHistoryFile(s string) error {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	errno := C.write_history(p)
	if errno == 0 && HistoryLength >= 0 {
		errno = C.history_truncate_file(p, C.int(HistoryLength))
	}
	if errno != 0 {
		return syscall.Errno(errno)
	}
	return nil
}

// SetCompleterDelims sets the word delimiters for tab-completion.
func SetCompleterDelims(break_chars string) {
	p := C.CString(break_chars)
	//defer C.free(unsafe.Pointer(p))
	C.free(unsafe.Pointer(C.rl_completer_word_break_characters))
	C.rl_completer_word_break_characters = p
}

// GetCompleterDemils gets current word delimiters for tab-completion.
func GetCompleterDelims() string {
	delims := C.GoString(C.rl_completer_word_break_characters)
	return delims
}

// CompletionMatches returns a list of completions for "text".
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

// SetAttemptedCompletionFunction sets the internal completion function to
// the custom function "cbk".
func SetAttemptedCompletionFunction(cbk func(text string, start, end int) []string) {
	c_cbk := (*C.rl_completion_func_t)(unsafe.Pointer(&cbk))
	C.rl_attempted_completion_function = c_cbk
}

/* EOF */
