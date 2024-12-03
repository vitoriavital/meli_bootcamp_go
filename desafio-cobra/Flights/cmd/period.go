package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var periodCmd = &cobra.Command{
    Use:     "period",
    Aliases: []string{"period"},
    Short:   "Get tickets of a period",
    Long:    "Get total flights tickets of 1 period",
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
		result,err := GetCountByPeriod(args[0])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Total tickets of %s period = %d tickets\n\n", args[0], result)
		}
    },
}

func init() {
    rootCmd.AddCommand(periodCmd)
}