package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App       *App
		DB        *DB
		Redis     *Redis
		Transport *Transport
	}
	App struct {
		Env  string
		Name string
	}
	DB struct {
		Connection string
		Host       string
		Port       string
		Name       string
		User       string
		Password   string
	}
	Redis struct {
		Host     string
		Port     string
		Password string
	}
	Transport struct {
		Env  string
		Host string
		Port string
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "prod" {
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
	}

	app := &App{
		Env:  os.Getenv("APP_ENV"),
		Name: os.Getenv("APP_NAME"),
	}
	redis := &Redis{
		Host:     os.Getenv("REDIS_ADDRESS"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}
	db := &DB{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		Name:       os.Getenv("DB_NAME"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
	}
	transport := &Transport{
		Env:  os.Getenv("APP_ENV"),
		Host: os.Getenv("TRANSPORT_ENV"),
		Port: os.Getenv("TRANSPORT_PORT"),
	}
	return &Container{
		App:       app,
		DB:        db,
		Redis:     redis,
		Transport: transport,
	}, nil
}
