package main

import (
	"os/exec"
	"os"
	"fmt"
)

func main(){
	/*cmd := exec.Command("vim")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
	*/
	cmd:=exec.Command("vim")
	fmt.Println(cmd.Dir)

}
