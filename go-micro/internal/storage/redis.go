package storage

import (
    "context"
    "time"

    "github.com/redis/go-redis/v9"
)

type RedisStore struct {
    client *redis.Client
}

func NewRedis(addr string) *RedisStore {
    rdb := redis.NewClient(&redis.Options{
        Addr: addr,
    })
    return &RedisStore{client: rdb}
}

func (r *RedisStore) SetOTP(ctx context.Context, phone, code string, ttl time.Duration) error {
    return r.client.Set(ctx, "otp:"+phone, code, ttl).Err()
}

func (r *RedisStore) GetOTP(ctx context.Context, phone string) (string, error) {
    return r.client.Get(ctx, "otp:"+phone).Result()
}

func (r *RedisStore) DeleteOTP(ctx context.Context, phone string) error {
    return r.client.Del(ctx, "otp:"+phone).Err()
}