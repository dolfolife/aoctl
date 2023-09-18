package aoc

import (
    "fmt"

    "github.com/spf13/cobra"
    "github.com/dolfolife/aoctl/pkg/aoc"
)
var puzzleCmd = &cobra.Command{
    Use: "puzzle",
    Aliases: []string{"p", "pzzl", "pz"},
    Short: "Get the puzzles for a given day",
    Args: cobra.ExactArgs(3),
    Run: func(cmd *cobra.Command, args []string) {
        day := args[0]
        year := args[1]
        cookie := args[2]
        puzzles := aoc.GetPuzzles(day, year, cookie)
        fmt.Println(puzzles[0])
        fmt.Println(puzzles[1])
    },
}

func init() {
    rootCmd.AddCommand(puzzleCmd)
}
