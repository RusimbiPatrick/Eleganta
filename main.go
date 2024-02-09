package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/RusimbiPatrick/Eleganta/repl"
)


func main(){
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is Eleganta version 0.01!\n", user.Username)
	fmt.Printf("Feel free to type in cmmands\n")
	repl.Start(os.Stdin, os.Stdout)
}