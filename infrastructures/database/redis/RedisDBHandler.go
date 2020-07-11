package redis

import (
	"errors"
	"sync"

	db "github.com/go-redis/redis"
)

// RedisDBHandler handles the methods for the redis database
type RedisDBHandler struct {
	m      sync.Mutex
	rw     sync.RWMutex
	Client *db.Client
}

// Connect to redis instance
func (r *RedisDBHandler) Connect(address string, password string, dbIndex int) (string, error) {
	r.m.Lock()
	defer r.m.Unlock()

	r.Client = db.NewClient(&db.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       dbIndex,  // use default DB
	})

	// basic test connection in go-redis
	pong, err := r.Client.Ping().Result()
	if err != nil {
		return "", err
	}

	return pong, err
}

// Flush flushes entire selected database
func (r *RedisDBHandler) Flush() {
	r.rw.Lock()
	defer r.rw.Unlock()

	r.Client.FlushDB()
}

// Delete a value by key
func (r *RedisDBHandler) Delete(key string) error {
	r.rw.Lock()
	defer r.rw.Unlock()

	err := r.Client.Del(key).Err()

	return err
}

// Get the value stored using a specified key (case-sensitive)
func (r *RedisDBHandler) Get(key string) (string, error) {
	r.rw.Lock()
	defer r.rw.Unlock()

	val, err := r.Client.Get(key).Result()
	if err == db.Nil {
		return "", errors.New("Empty")
	} else if err != nil {
		return "", err
	}

	return val, nil
}

// Set a key-value pair to redis database (case-sensitive)
func (r *RedisDBHandler) Set(key string, val interface{}) error {
	r.rw.Lock()
	defer r.rw.Unlock()

	err := r.Client.Set(key, val, 0).Err()

	return err
}
