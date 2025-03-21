package cmd

import (
	"log"
	"task/db"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		err := db.InitDB()
		if err != nil {
			log.Fatalf("Could not initialize database: %v", err)
		}

		defer db.CloseDB()

		db.ListTasks()
	},
}

func init() {
	RootCmd.AddCommand(ListCmd)
}
