package configs

import "flag"

var Server = new(ServerConfig)

type ServerConfig struct {
	Port int
}

func init() {
	flag.IntVar(&Server.Port, "server.port", 8080, "Server port")
	flag.Parse()
}
