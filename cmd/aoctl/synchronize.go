/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package aoctl

import (
	"github.com/dolfolife/aoctl/pkg/aoc"
	"github.com/spf13/cobra"
)

// syncronizeCmd represents the syncronize command
var (
	forced         = false
	synchronizeCmd = &cobra.Command{
		Use:   "synchronize",
		Short: "Synchronize current project",
		Long: `Synchronize allows you to update the local repository with the tree of problems in
adventofcode.com.`,
		Run: func(cmd *cobra.Command, args []string) {
			aoc.SyncAoC(forced)
		},
	}
)

func init() {
	puzzlesCmd.AddCommand(synchronizeCmd)

	synchronizeCmd.Flags().BoolVarP(&forced, "force", "f", false, "to force the files update")
}
