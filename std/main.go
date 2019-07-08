// You want to open the db everytime so that it can then run buckets in and out.
// There will be a bucket for each list which will be mapped out using the struct we have 
// read what we have into a file, open that file with text editor tool of some sort, edit and then read it back into bucket

package main

import (
	"github.com/spf13/cobra"
)
/* now what you can do with this is have this individual cobra commands use other cobra commands
*/


var rootCmd = &cobra.Command{ // this is a global variable dont do this! Put it in main instead
	// or do a function that returns the cobra command 
	Use: "", // will run everytime you type nothing in 
	Short: "Task Manager",
	Example: "An example would be that you enter in test[Enter] and then the command that you want to test",
}

func main(){
	rootCmd.AddCommand(createList)
	rootCmd.AddCommand(writeList)
	//rootCmd.AddCommand(deleteList)  // this will eventually be put into play 
	rootCmd.Execute()
}
