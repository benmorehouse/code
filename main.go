package main

import(
	"github.com/spf13/cobra"
	"fmt"
)

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
	rootCmd.AddCommand(runBMA)
	rootCmd.Execute()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err,"Error")
		os.Exit(1)
	}
}

var runBMA = &cobra.Command{
	Use: "run",
	Long: "Run your autofill editor tool",
	Short: "Run BME",
	Example: "./BME Run",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("runBMA")
	},
}

