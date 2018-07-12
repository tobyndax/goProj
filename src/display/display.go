package display

import "math/rand"
import "time"
import "fmt"

func PrintRandDelay(s string, maxDelayMs float64){
	for _, char := range s {
		fmt.Printf("%c",char)
		localDuration := time.Duration(maxDelayMs*rand.Float64())
		time.Sleep(localDuration*time.Millisecond)
	}
	fmt.Println("")
	return
}
