package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Sanjay-Aswar/Gophercises/task/cmd"
	"github.com/Sanjay-Aswar/Gophercises/task/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	cmd.RootCmd.Execute()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
