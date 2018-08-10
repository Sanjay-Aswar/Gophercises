package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Sanjay-Aswar/Gophercises/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Oops something gone wrong.")
			os.Exit(1)
		}
		fmt.Println("Successfully added task to task list with id :", id)

	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
