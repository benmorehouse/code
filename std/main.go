package main

import (
	"github.com/spf13/cobra"
	"fmt"
	/*"os"
	"github.com/boltdb/bolt"
	"log"
	"bufio"
	"strings"
	*/
)
/* now what you can do with this is have this individual cobra commands use other cobra commands
*/

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
	fmt.Println("Test:")
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(createListCmd)
	rootCmd.Execute()
	/*
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} // this is the initial test to see that the above var is okay 
	db, err := bolt.Open("database.db", 0600, nil) // this means that database now stands as a buffer for information
	var work list{
		head: nil,
		tail: nil,
		_size: 0,
	}

	if err != nil {
		log.Fatal(err) // the same as print and the OS.Exit(1)
	}
	fmt.Println("\n------------------------------------------------------------------")
	err = db.View(
	defer db.Close() // this is to ensure that regardless of what gappens in the code earlier, database closes
	*/
}

var listCmd = &cobra.Command{
	Use: "list",
	Long: "List items in the list of elements",
	Short: "Show items",
	Example: "./std list",
	Run: func(cmd *cobra.Command, args []string){
		for i:=0;i<len(args);i++{
			fmt.Println(args[i])
		}
		fmt.Println("Ran the list command")
		fmt.Println("test")
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

func loako(){
	for i:=0;i<10;i++{
		fmt.Print("4 loako")
	}
}

