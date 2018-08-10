package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark the task as complete from your task list",
	Run: func(cmd *cobra.Command, Args []string) {
		fmt.Println("do Called...")
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
