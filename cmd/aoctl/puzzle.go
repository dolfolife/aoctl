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
package aoctl

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dolfolife/aoctl/pkg/aoc"
	"github.com/spf13/cobra"
)

var puzzleCmd = &cobra.Command{
	Use:     "puzzle",
	Aliases: []string{"p", "pzzl", "pz"},
	Short:   "Get the puzzles for a given day",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		day := cmd.Flags().Lookup("day").Value.String()
		year := cmd.Flags().Lookup("year").Value.String()

		puzzles := aoc.GetPuzzles(day, year)
		fmt.Println(puzzles)
	},
}

func init() {
	rootCmd.AddCommand(puzzleCmd)

	puzzleCmd.Flags().StringP("day", "d", "1", "Day of the Advent of Code problems to fetch; By default it will be day 1")
	puzzleCmd.Flags().StringP("year", "y", getCurrentYear(), "Year of the Advent of Code; by default uses the last Advent of Code Year")
}

func getCurrentYear() string {
	year, month, _ := time.Now().Date()
	if month < 12 {
		return strconv.Itoa(year - 1)
	}

	return strconv.Itoa(year)
}
