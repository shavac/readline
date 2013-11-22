* ReadLine Wrapper for Golang *
===============================

This wrapper handlers SIGWINCH, and allows you to supply a Go-based
completion function.

--------------------------------------------------------------------------

    package main

    import "github.com/shavac/readline"

    func main() {
	    prompt := "by your command> ";
	    //loop until ReadLine returns nil (signalling EOF)
    L:
	    for {
		    switch result := readline.ReadLine(&prompt); true {
		    case result == nil:
		    	 println()
		    	 break L //exit loop with EOF(^D)
		    case *result != "": //ignore blank lines
			    println(*result);
			    readline.AddHistory(*result); //allow user to recall this line
		    }
	    }
    }

---------------------------------------------------------------------------
