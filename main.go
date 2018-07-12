package main
import "display"
import "interaction"

func main() {
	var delay float64 = 0
	display.PrintRandDelay("Hello... ",delay)
	display.PrintRandDelay("... can you hear me?",delay)
	display.PrintRandDelay("Confirm if the link has been established",delay)
	var input string = userInput.RequestAnyInput()
	
	display.PrintRandDelay("What do you mean by " + input,delay)

}
