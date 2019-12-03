package rsa

type ServerConfigParams struct {
	LogPath         string
	ConfigPath      string
	NodesConfigPath string
	DBPath          string
	Threshold       int
	Nodes           []string
}


func (conf *ServerConfigParams) Generate() error {

}