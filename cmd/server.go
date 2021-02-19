/*
@Time : 2021/2/8 13:33
@Author : wangtao
@File : run
*/
package cmd

import (
	"UDPexample/pkg/server"
	"flag"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run UDP Server",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println(strconv.Itoa(port))
		server.Server(ip, strconv.Itoa(port), mtu)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	serverCmd.Flags().StringVarP(&ip, "ip", "i", "0.0.0.0", "The IP that the UDP service listens on")
	serverCmd.Flags().IntVarP(&port, "port", "p", 8090, "The Port that the UDP service listens on")
	serverCmd.Flags().IntVarP(&mtu, "mtu", "m", 1500, "Read UDP msg len")
}
