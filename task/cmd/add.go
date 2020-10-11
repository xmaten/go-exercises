package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xmaten/task/db"
	"strings"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Adds a task to your list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")

		_, err := db.CreateTask(task)

		if err != nil {
			fmt.Println("Something went wrong: ", err.Error())
			return
		}

		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}