package display

import "math/rand"
import "time"
import "fmt"



func PrintRandDelay(s string, maxDelayMs float64){
	
	fmt.Print(">")
	
	for _, char := range s {
		fmt.Printf("%c",char)
		localDuration := time.Duration(maxDelayMs*rand.Float64())
		time.Sleep(localDuration*time.Millisecond)
	}
	time.Sleep(time.Duration(maxDelayMs)*time.Millisecond)
	fmt.Println("")
	time.Sleep(time.Duration(maxDelayMs)*time.Millisecond)

	return
}
