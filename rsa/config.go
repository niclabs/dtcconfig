package rsa

import (
	"fmt"
	dtc "github.com/niclabs/dtc/v2/config"
	node "github.com/niclabs/dtcnode/v2/config"
	"github.com/pebbe/zmq4"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"path"
	"strconv"
	"strings"
)

// ClientConfigParams groups the params received by Command Line
type ClientConfigParams struct {
	LogPath         string
	ConfigPath      string
	NodesConfigPath string
	DBPath          string
	Threshold       int
	Host            string
	Timeout         uint16
	Nodes           []string
}

// ClientConfig groups the three types of config used by DTC Client Library
type ClientConfig struct {
	LogFile string
	General dtc.Config
	Sqlite3 dtc.Sqlite3Config
	ZMQ     dtc.ZMQConfig
}

// GenerateConfig creates all the configuration related to RSA DTC implementation
func (conf *ClientConfigParams) GenerateConfig() error {
	if conf.Threshold > len(conf.Nodes) {
		return fmt.Errorf("threshold must be less or equal than nodes number")
	}
	servPubKey, servPrivKey, err := zmq4.NewCurveKeypair()
	if err != nil {
		return err
	}
	nodes, err := conf.CreateNodes(servPubKey)
	if err != nil {
		return err
	}
	// Create log file
	logDir := path.Dir(conf.LogPath)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return errors.Wrap(err, "Error creating directory for logging")
		}
	}
	if _, err := os.Stat(conf.LogPath); os.IsNotExist(err) {
		if _, err := os.OpenFile(conf.LogPath, os.O_CREATE, 0666); err != nil {
			return errors.Wrap(err, "Error creating logfile")
		}
	}
	clientConf := ClientConfig{
		LogFile: conf.LogPath,
		General: dtc.Config{
			DTC: dtc.DTCConfig{
				MessagingType: "zmq",
				NodesNumber:   uint16(len(conf.Nodes)),
				Threshold:     uint16(conf.Threshold),
			},
			Criptoki: dtc.CriptokiConfig{
				ManufacturerID:  "NICLabs",
				Model:           "dHSM RSA",
				Description:     "Distributed HSM using RSA signatures",
				SerialNumber:    "1",
				MinPinLength:    3,
				MaxPinLength:    10,
				MaxSessionCount: 5,
				DatabaseType:    "sqlite3",
				Slots: []*dtc.SlotsConfig{
					{
						Label: "TCBHSM",
						Pin:   "1234",
					},
				},
			},
		},
		Sqlite3: dtc.Sqlite3Config{
			Path: conf.DBPath,
		},
		ZMQ: dtc.ZMQConfig{
			PublicKey:  servPubKey,
			PrivateKey: servPrivKey,
			Nodes:      nodes,
			Timeout:    conf.Timeout,
		},
	}
	v := viper.New()
	v.Set("dtc", clientConf)
	configFolder := path.Dir(conf.ConfigPath)
	if _, err := os.Stat(configFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(configFolder, 0755); err != nil {
			return errors.Wrap(err, "Error creating config folder")
		}
	}
	if err := v.WriteConfigAs(conf.ConfigPath); err != nil {
		return errors.Wrap(err, "cannot write config file")
	}
	_, _ = fmt.Fprintf(os.Stderr, "config file written successfully in %s\n", conf.ConfigPath)
	return nil
}

//GenerateNodeConfig creates the configuration for the node files
func (conf *ClientConfigParams) GenerateNodeConfig(i int, clientPK string, nodeConf *dtc.NodeConfig, nodeSK string) error {
	outPath := path.Join(conf.NodesConfigPath, fmt.Sprintf("node_%d", i))
	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		if err := os.MkdirAll(outPath, 0755); err != nil {
			return errors.Wrap(err, fmt.Sprintf("cannot write node %d config folder", i))
		}
	}
	c := node.Config{
		PublicKey:  nodeConf.PublicKey,
		PrivateKey: nodeSK,
		Host:       nodeConf.Host,
		Port:       nodeConf.Port,
		Client: &node.ClientConfig{
			PublicKey: clientPK,
			Host:      conf.Host,
		},
	}
	v := viper.New()
	v.Set("config", c)
	if err := v.WriteConfigAs(path.Join(outPath, "config.yaml")); err != nil {
		return errors.Wrap(err, fmt.Sprintf("cannot write node %d config file", i))
	}
	_, _ = fmt.Fprintf(os.Stderr, "config file written successfully in %s\n", outPath)
	return nil
}

// CreateNodes creates
func (conf *ClientConfigParams) CreateNodes(clientPubKey string) ([]*dtc.NodeConfig, error) {
	dtcNodeConfig := make([]*dtc.NodeConfig, len(conf.Nodes))
	for i, aNode := range conf.Nodes {
		host, port, err := GetHostAndPort(aNode)
		if err != nil {
			return nil, err
		}
		nodePubKey, nodePrivKey, err := zmq4.NewCurveKeypair()
		if err != nil {
			return nil, err
		}
		dtcNodeConfig[i] = &dtc.NodeConfig{
			PublicKey: nodePubKey,
			Host:      host,
			Port:      port,
		}
		if err := conf.GenerateNodeConfig(i, clientPubKey, dtcNodeConfig[i], nodePrivKey); err != nil {
			return nil, err
		}
	}
	return dtcNodeConfig, nil
}

// GetHostAndPort splits a host and port string. Returns an error if something goes wrong.
func GetHostAndPort(ipPort string) (ip string, port uint16, err error) {
	nodeArr := strings.Split(ipPort, ":")
	if len(nodeArr) != 2 {
		err = fmt.Errorf("node ip and port format invalid. It should be ip:port\n")
		return
	}
	ip = nodeArr[0]
	portInt, err := strconv.Atoi(nodeArr[1])
	if err != nil {
		err = fmt.Errorf("could not convert port to int: %s\n", err)
		return
	}
	port = uint16(portInt)
	return
}
