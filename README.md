* ReadLine Wrapper for Golang *
===============================

Originally cloned from https://bitbucket.org/taruti/go-readline

Former code compile successfully with example but panic while window resizing.So i clone and patch it with
signal SIGWINCH handling with go codes.

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
