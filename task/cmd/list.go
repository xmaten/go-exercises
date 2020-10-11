package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xmaten/task/db"
	"os"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()

		if err != nil {
			fmt.Println("Something went wrong", err.Error())
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks")
			return
		}

		fmt.Println("Your tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i + 1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}