/*
@Time : 2021/2/8 13:41
@Author : wangtao
@File : root
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var ip string
var port int

var rootCmd = &cobra.Command{
	Use:   "UDPexample",
	Short: "a UDP Example",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
