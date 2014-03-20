// Copyright 2010-2014 go-readline authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

// This example demonstrates how to create a simple readline application
// that allows you to tab-complete single letters with the various words
// used in phonetic alphabets. See:
// http://usmilitary.about.com/od/theorderlyroom/a/alphabet.htm

import (
	"fmt"
	"strings"

	"github.com/fiorix/go-readline"
)

var phoneticAlphabet = map[string][]string{
	"a": {"Alpha", "Able", "Affirmative", "Afirm"},
	"b": {"Bravo", "Boy", "Baker"},
	"c": {"Charlie", "Cast"},
	"d": {"Delta", "Dog"},
	"e": {"Echo", "Easy"},
	"f": {"Foxtrot", "Fox"},
	"g": {"Golf", "George"},
	"h": {"Hotel", "Have", "Hypo", "How"},
	"i": {"India", "Item", "Interrogatory", "Int", "Item"},
	"j": {"Juliett", "Jig"},
	"k": {"Kilo", "King"},
	"l": {"Lime", "Love"},
	"m": {"Mike"},
	"n": {"November", "Nan", "Negative", "Negat"},
	"o": {"Oscar", "Oboe", "Option", "Oboe"},
	"p": {"Papa", "Pup", "Preparatory", "Prep", "Peter"},
	"q": {"Quebec", "Quack", "Queen"},
	"r": {"Romeo", "Rush", "Roger"},
	"s": {"Sierra", "Sail", "Sugar"},
	"t": {"Tango", "Tare"},
	"u": {"Uniform", "Unit", "Uncle"},
	"v": {"Victor", "Vice"},
	"w": {"Whiskey", "Watch", "William"},
	"x": {"X-ray"},
	"y": {"Yankee", "Yoke"},
	"z": {"Zulu", "Zed", "Zebra"},
}

func completer(input, line string, start, end int) []string {
	if len(input) == 1 {
		letters, exists := phoneticAlphabet[strings.ToLower(input)]
		if exists {
			return letters
		}
	}
	return []string{}
}

func main() {
	prompt := "> "

	readline.SetCompletionFunction(completer)

	// This is generally what people expect in a modern Readline-based app
	readline.ParseAndBind("TAB: menu-complete")

	// Loop until Readline returns nil (signalling EOF)
L:
	for {
		result := readline.Readline(&prompt)
		switch {
		case result == nil:
			fmt.Println()
			break L // exit loop
		case *result == "exit":
			break L // exit loop
		case *result != "": // Ignore blank lines
			fmt.Println(*result)
			readline.AddHistory(*result) // Allow user to recall this line
		}
	}
}
