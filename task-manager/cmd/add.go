package cmd

import (
	"fmt"
	"log"
	"task/db"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a task description.")
			return
		}

		// Initialize the database
		err := db.InitDB()
		if err != nil {
			log.Fatalf("Could not initialize database: %v", err)
		}
		defer db.CloseDB()

		task := args[0]
		err = db.AddTask(task)
		if err != nil {
			log.Fatalf("Could not add task: %v", err)
		}
		fmt.Printf("Added task: %s\n", task)
	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}
