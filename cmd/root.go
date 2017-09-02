// Copyright © 2017 Skitta Chen <skittachen@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/skitta/gopinyin/pinyin"
)

var (
	heteronym bool
	style     string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gopinyin",
	Short: "translate a Chinese charecter into real toun pinyin",
	Long:  `输入汉字，得到拼音`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: gopinyin,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.Flags().BoolVarP(&heteronym, "heteronym", "e", false, "启用多音字模式")
	RootCmd.Flags().StringVarP(&style, "style", "s", "Tone", "指定拼音风格。可选值：Normal, Tone, Tone2, Tone3, Initials, FirstLetter, Finals, FinalsTone, FinalsTone2, FinalsTone3")
}

func gopinyin(cmd *cobra.Command, args []string) error {
	pargs := pinyin.NewArgs()

	if heteronym {
		pargs.Heteronym = true
	}

	switch style {
	case "Nomal":
		pargs.Style = pinyin.Normal
	case "Tone2":
		pargs.Style = pinyin.Tone2
	case "Tone3":
		pargs.Style = pinyin.Tone3
	case "Initials":
		pargs.Style = pinyin.Initials
	case "FirstLetter":
		pargs.Style = pinyin.FirstLetter
	case "Finals":
		pargs.Style = pinyin.Finals
	case "FinalsTone":
		pargs.Style = pinyin.FinalsTone
	case "FinalsTone2":
		pargs.Style = pinyin.FinalsTone2
	case "FinalsTone3":
		pargs.Style = pinyin.FinalsTone3
	default:
		pargs.Style = pinyin.Tone
	}

	pys := pinyin.Pinyin(strings.Join(args, ""), pargs)
	for _, s := range pys {
		fmt.Print(strings.Join(s, ","), " ")
	}
	if len(pys) > 0 {
		fmt.Println()
	}

	return nil
}
