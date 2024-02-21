package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName      string `env:"APP_NAME" envDefault:"Auth_Service"`
	HostGRPC     string
	PortGRPC     int
	UsernameGRPC string
	PasswordGRPC string
	HostDB       string
	PortDB       int
	LoginDB      string
	PasswordDB   string
	NameDB       string
	SSLTypeDB    string
}

func NewConfig() (*Config, error) {
	var cfg Config

	err := godotenv.Load(".env." + "local")
	if err != nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Ошибка при загрузке файла .env: - ", err)
		}
	}

	cfg.AppName = os.Getenv("APP_NAME")

	cfg.HostGRPC = os.Getenv("GRPC_HOST")
	cfg.PortGRPC, err = strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("Ошибка при преобразовании порта GRPC в число: %v", err)
	}
	cfg.UsernameGRPC = os.Getenv("GRPC_LOGIN")
	cfg.PasswordGRPC = os.Getenv("GRPC_PASSWORD")

	cfg.HostDB = os.Getenv("POSTGRES_HOST")
	if len(cfg.HostDB) < 1 {
		cfg.HostDB = "localhost"
	}
	cfg.PortDB, err = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Fatalf("Ошибка при преобразовании порта базы данных в число: %v", err)
	}
	cfg.LoginDB = os.Getenv("POSTGRES_USER")
	if len(cfg.LoginDB) < 1 {
		log.Fatalf("POSTGRES_USER - empty")
	}
	cfg.PasswordDB = os.Getenv("POSTGRES_PASSWORD")
	cfg.NameDB = os.Getenv("POSTGRES_DB")
	cfg.SSLTypeDB = os.Getenv("POSTGRES_SSL")
	if len(cfg.SSLTypeDB) < 1 {
		cfg.SSLTypeDB = "disable"
	}

	return &cfg, nil
}
