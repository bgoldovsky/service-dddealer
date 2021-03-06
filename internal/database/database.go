package database

import (
	"context"
	"net"
	"time"

	"github.com/bgoldovsky/dealer/service-dealer/internal/app/logger"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDatabase(ctx context.Context, connString string) *pgxpool.Pool {
	c, err := pgxpool.ParseConfig(connString)
	if err != nil {
		panic("failed to parse postgres config: " + err.Error())
	}

	c.MaxConns = 50
	c.MinConns = 10
	c.ConnConfig.TLSConfig = nil
	c.ConnConfig.PreferSimpleProtocol = true
	c.ConnConfig.RuntimeParams["standard_conforming_strings"] = "on"
	c.ConnConfig.DialFunc = (&net.Dialer{
		KeepAlive: 5 * time.Minute,
		Timeout:   1 * time.Second,
	}).DialContext

	p, err := pgxpool.ConnectConfig(ctx, c)
	if err != nil {
		panic("failed to connect to postgres: " + err.Error())
	}

	logger.Log.WithField("cs", connString).Info("CONNECTED")
	return p
}
