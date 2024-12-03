package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var averageCmd = &cobra.Command{
    Use:     "average",
    Aliases: []string{"average"},
    Short:   "Get average percentage of tickets to destination",
    Long:    "Get average percentage of tickets to 1 destination",
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
		result,err := AverageDestination(args[0], 1000)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Average percentage of tickets to %s = %0.01f%%\n\n", args[0], result)
		}
    },
}

func init() {
    rootCmd.AddCommand(averageCmd)
}