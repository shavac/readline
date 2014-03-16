package main

import "github.com/shavac/readline"

func main() {
	prompt := "> "
	//loop until ReadLine returns nil (signalling EOF)
L:
	for {
		switch result := readline.ReadLine(&prompt); true {
		case result == nil:
			println()
			break L //exit loop
		case *result != "": //ignore blank lines
			println(*result)
			readline.AddHistory(*result) //allow user to recall this line
		}
	}
}
