package ent

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go_template/conf"
)

func NewEntClient(lc fx.Lifecycle, log *zap.Logger, conf *conf.Config) *Client {
	client, err := Open("sqlite3", conf.DBDsn)
	if err != nil {
		log.Error("failed opening connection to sqlite3", zap.Error(err))
		panic(err)
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			client.Close()
			log.Info("ent client is closed")
			return nil
		},
	})
	return client

}
