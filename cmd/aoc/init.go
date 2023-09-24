/*
Copyright © 2023 Rodolfo Sanchez <rodolfo2488@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package aoc

import (

    "github.com/spf13/cobra"
    "github.com/dolfolife/aoctl/pkg/aoc"
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
