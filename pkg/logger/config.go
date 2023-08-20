package logger

type Config struct {
	Development bool   `koanf:"developmet"`
	Encoding    string `koanf:"encoding"`
	Level       string `koanf:"level"`
}
