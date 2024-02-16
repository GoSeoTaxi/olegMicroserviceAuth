package config

const grpcPort = 50051

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewConfig() *Config {
	return &Config{
		Host:     "localhost",
		Port:     grpcPort,
		Username: "",
		Password: "",
	}
}
