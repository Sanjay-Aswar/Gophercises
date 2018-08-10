package cmd

import (
	"fmt"
	"strconv"

	"github.com/Sanjay-Aswar/Gophercises/task/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark the task as complete from your task list",
	Run: func(cmd *cobra.Command, args []string) {
		for _, i := range args {
			pref, err := strconv.Atoi(i)
			if err != nil {
				fmt.Printf("Error in parsing %s preference", i)
			}
			db.DeleteTask(pref)
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
