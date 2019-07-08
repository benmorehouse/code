package main

import (
	"os"
	"os/exec"
)

func main(){
	cmd := exec.Command("vim","-o","temp.txt")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}
