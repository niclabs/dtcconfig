package cmd

import (
	"github.com/niclabs/dtcconfig/rsa"
	"github.com/spf13/cobra"
)

var serverConfig rsa.ClientConfigParams

func init() {
	rsaCmd.Flags().StringVarP(
		&serverConfig.Host,
		"host",
		"H",
		"",
		"IP or domain name that the nodes will see from the client")
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
		"./config.yaml",
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
		"/tmp/dtc.log",
		"path to a file where to output the services logs")
	rsaCmd.Flags().StringVarP(
		&serverConfig.DBPath,
		"db",
		"d",
		"./db.sqlite3",
		"path to a file where to put Sqlite3 database")
	rsaCmd.Flags().IntVarP(
		&serverConfig.Threshold,
		"threshold",
		"t",
		0,
		"Minimum number of nodes required to sign")
	_ = rsaCmd.MarkFlagRequired("host")
	_ = rsaCmd.MarkFlagRequired("nodes")
	_ = rsaCmd.MarkFlagRequired("threshold")
}

var rsaCmd = &cobra.Command{
	Use:   "rsa",
	Short: "Generates a configuration file for tcrsa nodes and server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return serverConfig.GenerateConfig()
	},
}
