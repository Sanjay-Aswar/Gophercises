package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.ibm.com/sanjay.aswar1/gophercises/task/db"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks from your task list",
	Run: func(cmd *cobra.Command, Args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Somthing went wrong.")
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You don't have any tasks to complete.")
			return
		}
		fmt.Println("You have following tasks to complete.")
		for i, task := range tasks {
			fmt.Printf("%d\t%v", i, task)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
