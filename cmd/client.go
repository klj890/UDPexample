/*
@Time : 2021/2/8 15:00
@Author : wangtao
@File : client
*/
package cmd

import (
	"UDPexample/pkg/client"
	"flag"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"strconv"
)

var dataPath string

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run a UDP Client",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		client.Client(ip, strconv.Itoa(port), dataPath, mtu)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	clientCmd.Flags().StringVarP(&ip, "ip", "i", "127.0.0.1", "Destination IP")
	clientCmd.Flags().IntVarP(&port, "port", "p", 8090, "Destination port")
	clientCmd.Flags().StringVarP(&dataPath, "data", "d", "./examples/data1.json", "Send data file path")
	clientCmd.Flags().IntVarP(&mtu, "mtu", "m", 1500, "Send UDP msg length")
}
