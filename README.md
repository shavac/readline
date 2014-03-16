# go-readline

go-readline is a wrapper for the
[GNU readline library](http://cnswww.cns.cwru.edu/php/chet/readline/rltop.html)
for the [Go programming language](http://golang.org).

This repository contains work from multiple contributors. See the AUTHORS file
for details.




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
