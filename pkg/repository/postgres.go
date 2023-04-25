package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SslMode  string
}

func NewPostgresDB(c Config) (*sqlx.DB, error) { // Подключение к базе данных
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.Username, c.Password, c.DBName, c.SslMode,
	))
	if err != nil { // если ошибка в подключении к бд есть то мы возвращаем ошибку
		return nil, err
	}
	err = db.Ping() // пингуем, чтобы проверить, что подключение есть
	if err != nil { // если при пинге подключения к бд нет, то мы возвращаем ошибку
		return nil, err
	}
	return db, nil // возвращаем бд
}
