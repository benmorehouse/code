package main

import (
	"os/exec"
	"os"
)

func main(){
	cmd := exec.Command("pwd")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
	cmd = exec.Command("pwd")
	cmd.Run()
}


// use this to go into vim and just edit around the line 9 command 
