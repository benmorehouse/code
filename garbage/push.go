package main

import (
	"os/exec"
	"os"
)

// Example of how to open a repository in a specific path, and push to
// its default remote (origin).
func main() {

/*
	CheckArgs("<repository-path>")
	path := os.Args[1]
	fmt.Println(path)
	r, err := git.PlainOpen(path)
	CheckIfError(err)

	Info("git push")
	// push using default options
	err = r.Push(&git.PushOptions{})
	CheckIfError(err)
*/
	openFile()
}

func openFile(){
	cmd := exec.Command("git push")
	// the hope is that buffer.txt will have everything loaded in from the bucket
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}


