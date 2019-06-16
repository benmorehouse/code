package main

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Long: "If you wanna use this test command, then enter in some sort of command that you see",
	Short: "test is just gonna output test the command it prompts you with",
	Example: "An example would be that you enter in test[Enter] and then the command that you want to test",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func main(){
	if err := testCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} // this is the initial test to see that the above var is okay 
}
