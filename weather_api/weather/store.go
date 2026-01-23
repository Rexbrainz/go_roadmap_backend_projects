package weather

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type MemStore struct {
	Cache	map[string]Weather
}

type RedisStore struct {
	client	*redis.Client
}

// A database interface for caching.
type Storer interface {
	Get(string) (Weather, bool)
	Set(string, Weather, time.Duration)
}

// Get a weather report if cached
func (m *MemStore) Get(city string) (Weather, bool) {
	report, ok := m.Cache[city]
	if !ok {
		return Weather{}, false
	}
	return report, true
}

// Cache a weather report
func (m *MemStore) Set(city string, weather Weather, ttl time.Duration) {
	m.Cache[city] = weather
}

// Construct and return an in memory store
func NewMemStore() *MemStore {
	return &MemStore{
		Cache: map[string]Weather{},
	}
}

// Construct and return a Redis Store
func NewRedisStore(redisURL string) (*RedisStore, error) {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &RedisStore{
		client: client,
	}, nil
}

// Get a weather report if cached
func (r *RedisStore) Get(city string) (Weather, bool) {
	ctx := context.Background()
	key := redisKey(city)

	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return Weather{}, false 
	}
	if err != nil {
		return Weather{}, false
	}

	var w Weather
	if err := json.Unmarshal([]byte(val), &w); err != nil {
		return Weather{}, false
	}
	return w, true
}

// Cache a weather report
func(r *RedisStore) Set(city string, weather Weather, ttl time.Duration) {
	ctx := context.Background()
	key := redisKey(city)

	data, err := json.Marshal(weather)
	if err != nil {
		return
	}

	// Redis is an optimization, if it fails the weather server should still work
	_ = r.client.Set(ctx, key, data, ttl).Err()
}

func redisKey(city string) string {
	return "weather:" + city
}