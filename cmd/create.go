package cmd

import (
	"github.com/niclabs/dtcconfig/config"
	"github.com/spf13/cobra"
)

var serverConfig config.ClientConfigParams

func init() {
	createCmd.Flags().StringVarP(
		&serverConfig.Host,
		"host",
		"H",
		"",
		"IP or domain name that the nodes will see from the client")
	createCmd.Flags().StringSliceVarP(
		&serverConfig.Nodes,
		"nodes",
		"n",
		[]string{},
		"comma separated list of nodes in ip:port format")
	createCmd.Flags().StringVarP(
		&serverConfig.ConfigPath,
		"config",
		"c",
		"./dtc-config.yaml",
		"path where to output the local config file")
	createCmd.Flags().StringVarP(
		&serverConfig.NodesConfigPath,
		"nodes-config",
		"k",
		"./nodes",
		"path to a folder where to output the nodes config files")
	createCmd.Flags().StringVarP(
		&serverConfig.LogPath,
		"log",
		"l",
		"/tmp/dtc.log",
		"path to a file where to output the services logs")
	createCmd.Flags().StringVarP(
		&serverConfig.DBPath,
		"db",
		"d",
		"./db.sqlite3",
		"path to a file where to put Sqlite3 database")
	createCmd.Flags().IntVarP(
		&serverConfig.Threshold,
		"threshold",
		"t",
		0,
		"Minimum number of nodes required to sign")
	_ = createCmd.MarkFlagRequired("host")
	_ = createCmd.MarkFlagRequired("nodes")
	_ = createCmd.MarkFlagRequired("threshold")
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Generates a configuration file for dtc nodes and server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return serverConfig.GenerateConfig()
	},
}
