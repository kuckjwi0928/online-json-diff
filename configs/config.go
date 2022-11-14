package configs

import "flag"

var Server = new(ServerConfig)

type ServerConfig struct {
	Port int
	Env  string
}

func init() {
	flag.IntVar(&Server.Port, "server.port", 8080, "Server port")
	flag.StringVar(&Server.Env, "server.env", "dev", "Server environment")
	flag.Parse()
}
