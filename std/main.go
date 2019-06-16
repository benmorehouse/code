package main

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)
/* now what you can do with this is have this individual cobra commands use other cobra commands
*/

func loako(){
	for i:=0;i<10;i++{
		fmt.Print("loako ")
	}
}

var rootCmd = &cobra.Command{
	Use: "root",
	Long: "If you wanna use this test command, then enter in some sort of command that you see",
	Short: "test is just gonna output test the command it prompts you with",
	Example: "An example would be that you enter in test[Enter] and then the command that you want to test",
	Run: func(cmd *cobra.Command, args []string) {
		// when ever you run STD this always runs. Maybe you can get this to show you all of the lists
	},
}

func main(){
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(createListCmd)
	rootCmd.AddCommand(helpCmd)
	rootCmd.Execute()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} // this is the initial test to see that the above var is okay 
}

var listCmd = &cobra.Command{
	Use: "list",
	Long: "List items in the list of elements",
	Short: "Show items",
	Example: "./std list",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("Ran the list command")
	},
}

var createListCmd = &cobra.Command{
	Use: "create",
	Long: "Creats a list of things to do",
	Short: "create list",
	Example: "./std create",
	Run: func(cmd *cobra.Command, args []string){
		loako()
	},
}

