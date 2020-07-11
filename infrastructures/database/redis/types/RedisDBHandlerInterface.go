package types

// RedisDBHandlerInterface holds the list of implementable methods for the RedisDBHandler
type RedisDBHandlerInterface interface {
	Connect(address string, password string, dbIndex int) (string, error)
	Flush()
	Delete(key string) error
	Get(key string) (string, error)
	Set(key string, val interface{}) error
}
