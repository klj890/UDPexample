/*
@Time : 2021/2/19 15:16
@Author : wangtao
@File : multicast
*/
package cmd

import (
	"UDPexample/pkg/server"
	"flag"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var multiSvc = &cobra.Command{
	Use:   "multicast",
	Short: "Run Multicast UDP Server",
	Run: func(cmd *cobra.Command, args []string) {
		server.MultServer(ip, port, mtu)
	},
}

func init() {
	rootCmd.AddCommand(multiSvc)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	multiSvc.Flags().StringVarP(&ip, "ip", "i", "0.0.0.0", "The IP that the UDP service listens on")
	multiSvc.Flags().IntVarP(&port, "port", "p", 8090, "The Port that the UDP service listens on")
	multiSvc.Flags().IntVarP(&mtu, "mtu", "m", 1500, "Read UDP msg len")
}
