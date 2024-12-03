package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var totalCmd = &cobra.Command{
    Use:     "total",
    Aliases: []string{"total"},
    Short:   "Get total tickets",
    Long:    "Get total flights tickets of 1 destination",
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
		result,err := GetTotalTickets(args[0])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Total flight tickets to: %s = %d tickets\n\n", args[0], result)
		}
    },
}

func init() {
    rootCmd.AddCommand(totalCmd)
}