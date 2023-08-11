package aoc

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "aoc",
    Short: "Advent of Code CLI tool",
    Long:  `Advent of Code CLI tool to download and run solutions`,
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "There was an error while executing the command %s", err)
        os.Exit(1)
    }
}
