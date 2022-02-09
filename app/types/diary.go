package types

type LocalInfo struct {
	Db_host string
	Db_port string
	Db_name string
	Db_username string
	Db_password string
}

type Server struct {
	Server int `toml:"port"`
	Local LocalInfo `toml:"local"`
}

type Config struct {
	Server     Server
	ConfigPath string
}