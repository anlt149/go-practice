package cmd

import (
	"log"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var DoCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your task list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("Please provide a task ID.")
			return
		}

		err := db.InitDB()

		if err != nil {
			log.Fatalf("Could not initialize database: %v", err)
		}

		defer db.CloseDB()
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Invalid task ID: %v", err)
		}
		db.UpdateTask(taskID)
	},
}

func init() {
	RootCmd.AddCommand(DoCmd)
}
