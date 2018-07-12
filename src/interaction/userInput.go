package userInput

import "fmt"
import "bufio"
import "strings"
import "os"
import "log"
import "path/filepath"



type PlayerInfo struct{
	Name string
	Username string
}

var Info PlayerInfo
//Special function init, will run upon package import. 

func init(){
	Info.Username = os.Getenv("USER")
	Info.Name = os.Getenv("NAME")
}

func emulateConsole(){
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	dir = strings.TrimPrefix(dir, "/home/" + Info.Username)
	dir = Info.Username + "@" + Info.Name + ":" + "~" + dir + "$ "
	fmt.Print(dir)
}

func RequestAnyInput() string{
	emulateConsole()
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\n","",-1)
	return text
}
