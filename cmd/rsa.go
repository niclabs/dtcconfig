package cmd

import (
	"github.com/niclabs/dtcconfig/rsa"
	"github.com/spf13/cobra"
)

var serverConfig rsa.ServerConfigParams

func init() {
	rsaCmd.Flags().StringSliceVarP(
		&serverConfig.Nodes,
		"nodes",
		"n",
		[]string{},
		"comma separated list of nodes in ip:port format")
	rsaCmd.Flags().StringVarP(
		&serverConfig.ConfigPath,
		"config",
		"c",
		"/etc/dtc/config.yaml",
		"path where to output the local config file")
	rsaCmd.Flags().StringVarP(
		&serverConfig.NodesConfigPath,
		"nodes-config",
		"k",
		"./nodes",
		"path to a folder where to output the nodes config files")
	rsaCmd.Flags().StringVarP(
		&serverConfig.LogPath,
		"log",
		"l",
		"/var/log/dtc.log",
		"path to a file where to output the services logs")
	rsaCmd.Flags().StringVarP(
		&serverConfig.DBPath,
		"db",
		"d",
		"/etc/dtc/db.sqlite3",
		"path to a file where to put Sqlite3 database")
	rsaCmd.Flags().IntVarP(
		&serverConfig.Threshold,
		"threshold",
		"t",
		len(serverConfig.Nodes)/2,
		"Minimum number of nodes required to sign")
	_ = rsaCmd.MarkFlagRequired("nodes")
	_ = rsaCmd.MarkFlagRequired("local-config")
	_ = rsaCmd.MarkFlagRequired("nodes-config")
}

var rsaCmd = &cobra.Command{
	Use:   "rsaCmd",
	Short: "Generates a configuration file for tcrsa nodes and server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return serverConfig.Generate()
	},
}
