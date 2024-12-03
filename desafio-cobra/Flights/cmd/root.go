package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "flights",
    Short: "flights is a cli tool for basic flights stats",
    Long:  "flights is a cli tool  for basic flights stats - GetTotalTickets, GetCountByPeriod, AverageDestination.",
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Oops. An error while executing Flights '%s'\n", err)
        os.Exit(1)
    }
}