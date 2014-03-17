// Copyright 2010-2014 go-readline authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import "github.com/fiorix/go-readline"

func main() {
	prompt := "> "

	// Loop until Readline returns nil (signalling EOF)
L:
	for {
		result := readline.Readline(&prompt)
		switch {
		case result == nil:
			println()
			break L // exit loop
		case *result != "": // Ignore blank lines
			println(*result)
			readline.AddHistory(*result) // Allow user to recall this line
		}
	}
}
