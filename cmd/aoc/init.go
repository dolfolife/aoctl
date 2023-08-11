package aoc

import (

    "github.com/spf13/cobra"
    "github.com/dolfolife/adventofcode/pkg/aoc"
)

var initCmd = &cobra.Command{
    Use: "init",
    Aliases: []string{"initialize"},
    Short: "Initialize the Advent of Code project in the path specifed",
    Args: cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
        path := cmd.Flags().Lookup("path").Value.String()

        if path == "" {
            path= "adventofcode"
        }

        aoc.InitializeProject(path)
    },
}

func init() {
    rootCmd.AddCommand(initCmd)

    initCmd.Flags().StringP("path", "p", "", "Path to initialize the project")
}
