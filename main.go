package main //package tells application to get ready

import (
	"fmt"
) // if line 6 fail type import (fmt) like here

func main() { //main entry point, this is where the code runs and displays on the terminal
	/*|-> if you create a new file in the new folder, main should change to the name of file
	to avoid duplicate directory*/

	fmt.Print("Hello, world") //fmt tells 'go' to format the text by printing it

}

//this is a standard coding skeleton
