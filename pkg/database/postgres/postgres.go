package postgres

import (
	"SystemOfAuthorization/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"time"
)

func StartPostgres(cfg config.Config, log *logrus.Logger) (*pgxpool.Pool, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s&sslmode=%s",
		cfg.DBUser, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Infoln("Ошибка подключения в строке к БД:", err)
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		log.Error("Ошибка подключения к бд")
		return nil, err
	}

	return conn, nil

}
