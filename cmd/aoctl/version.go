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

    goVersion "go.hein.dev/go-version"
	"github.com/spf13/cobra"
)

// versionCmd represents the aoc/version command
var (
    shortened  = false
	version    = "dev"
	commit     = "none"
	date       = "unknown"
	output     = "json"
    versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version of the Advent of Code CLI",
	Run: func(cmd *cobra.Command, args []string) {
            fmt.Printf("version: %s\n", version)
            var response string
	        versionOutput := goVersion.New(version, commit, date)

	        if shortened {
                response = versionOutput.ToShortened()
            } else {
                response = versionOutput.ToJSON()
            }
            fmt.Printf("%+v", response)
            return
	    },
    }
)

func init() {
    versionCmd.Flags().BoolVarP(&shortened, "short", "s", false, "Print just the version number.")
	versionCmd.Flags().StringVarP(&output, "output", "o", "json", "Output format. One of 'yaml' or 'json'.")

	rootCmd.AddCommand(versionCmd)
}
