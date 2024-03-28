package config

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

func NewPGConfig() (PGConfig, error) {

	var host, portin, login, pass, nameDB, sslType, dsn string

	host = os.Getenv("POSTGRES_HOST")
	if len(host) < 1 {
		host = "localhost"
	}
	portin = os.Getenv("POSTGRES_PORT")

	if len(portin) < 1 {
		portin = "15432"
	}

	/*	porti, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
		if err != nil {
			log.Fatalf("Ошибка при преобразовании порта базы данных в число: %v", err)
		}*/
	login = os.Getenv("POSTGRES_USER")
	if len(login) < 1 {
		log.Fatalf("POSTGRES_USER - empty")
	}
	pass = os.Getenv("POSTGRES_PASSWORD")
	nameDB = os.Getenv("POSTGRES_DB")
	sslType = os.Getenv("POSTGRES_SSL")
	if len(sslType) < 1 {
		sslType = "disable"
	}

	dsn = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v\n", host, portin, login, pass, nameDB, sslType)

	if len(dsn) == 0 {
		return nil, errors.New("pg dsn not found")
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
