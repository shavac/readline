// test program for the readline package
package main

import "readline"

func main() {
	prompt := "by your command> ";

	//loop until ReadLine returns nil (signalling EOF)
	L: for {
		switch result := readline.ReadLine(&prompt); true {
		case result == nil: break L //exit loop

		case *result != "": //ignore blank lines
			println(*result);
			readline.AddHistory(*result); //allow user to recall this line
		}
	}
}
