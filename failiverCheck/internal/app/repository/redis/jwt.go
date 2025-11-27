package redis

import (
	"context"
	"time"
)

const jwtPrefix = "jwt:"

func (c *Redis) getJwtKey(token string) string {
	return c.cfg.AppPrefix + jwtPrefix + token
}

func (c *Redis) SetBlackListJWT(ctx context.Context, token string, jwtTTL time.Duration) error {
	return c.client.Set(ctx, c.getJwtKey(token), true, jwtTTL).Err()
}

func (c *Redis) GetBlackListJWT(ctx context.Context, token string) error {
	return c.client.Get(ctx, c.getJwtKey(token)).Err()
}
