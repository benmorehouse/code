package main

import (
	"os/exec"
	"os"
	"fmt"
)

func main() {
	err := sendText()
	fmt.Println(err)
}

func sendText()error{
	cmd := exec.Command("osascript","sendMessage.applescript","+12064581383", "sup")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}
