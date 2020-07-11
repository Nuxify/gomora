package redis

import (
	"testing"
)

func TestConnection(t *testing.T) {
	db := new(RedisDBHandler)

	res, err := db.Connect("localhost:6379", "", 0)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(res)
}

func TestWriteRead(t *testing.T) {
	t.Run("write some data", func(t *testing.T) {
		// connect to redis
		db := new(RedisDBHandler)

		_, err := db.Connect("localhost:6379", "", 0)
		if err != nil {
			t.Error(err)
			return
		}

		// write
		err = db.Set("erply", "please accept me :D")
		if err != nil {
			t.Error(err)
			return
		}

		t.Log("success")
		return
	})

	t.Run("read some data", func(t *testing.T) {
		// connect to redis
		db := new(RedisDBHandler)

		_, err := db.Connect("localhost:6379", "", 0)
		if err != nil {
			t.Error(err)
			return
		}

		// read
		data, err := db.Get("erply")
		if err != nil {
			t.Error(err)
			return
		}

		t.Log(data)
		return
	})
}

func TestDeleteByKey(t *testing.T) {
	// connect to redis
	db := new(RedisDBHandler)

	_, err := db.Connect("localhost:6379", "", 0)
	if err != nil {
		t.Error(err)
		return
	}

	err = db.Delete("erply")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("successfully deleted")
	return
}

func TestFlushDB(t *testing.T) {
	// connect to redis
	db := new(RedisDBHandler)

	_, err := db.Connect("localhost:6379", "", 0)
	if err != nil {
		t.Error(err)
		return
	}

	db.Flush()
}
