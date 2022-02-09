package types

type Server struct {
	Server int `toml:"port"`
	Local LocalInfo `toml:"local"`
}