package main
import "fmt"
import "math/rand"
import "time"
import "display"

func printRandDelay(s string, maxDelayMs float64){
	for _, char := range s {
		fmt.Printf("%c",char)
		localDuration := time.Duration(maxDelayMs*rand.Float64())
		time.Sleep(localDuration*time.Millisecond)
	}
	fmt.Println("")
	return
}

func main() {
	display.NewFunc()
	printRandDelay("Hello World!",200)
}