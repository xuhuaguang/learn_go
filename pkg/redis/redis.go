package redis

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"time"
)

var NotFoundErr = redis.ErrNil

const (
	REDIS_KEYWORD_HSET    = "HSET"
	REDIS_KEYWORD_HGET    = "HGET"
	REDIS_KEYWORD_HMSET   = "HMSET"
	REDIS_KEYWORD_HMGET   = "HMGET"
	REDIS_KEYWORD_HDEL    = "HDEL"
	REDIS_KEYWORD_HGETALL = "HGETALL"

	REDIS_KEYWORD_INCR        = "INCR"
	REDIS_KEYWORD_DECR        = "DECR"
	REDIS_KEYWORD_INCRBY      = "INCRBY"
	REDIS_KEYWORD_DECRBY      = "DECRBY"
	REDIS_KEYWORD_INCRBYFLOAT = "INCRBYFLOAT"
	REDIS_KEYWORD_DECRBYFLOAT = "DECRBYFLOAT"
	REDIS_KEYWORD_HINCRBY     = "HINCRBY"
)

type ConnPool struct {
	raw *redis.Pool
}

func NewPool(opt *PoolOption) (*ConnPool, error) {
	rawPool := newPool(opt)
	if rawPool == nil {
		return nil, errors.New("init redis raw failed")
	}
	// 测试连接池是否正常
	conn := rawPool.Get()
	defer conn.Close()
	_, err := conn.Do("PING")
	if err != nil || conn.Err() != nil {
		return nil, conn.Err()
	}
	return &ConnPool{rawPool}, nil
}

func newPool(opt *PoolOption) *redis.Pool {
	return &redis.Pool{
		MaxActive:   opt.MaxActive,
		MaxIdle:     opt.MaxIdle,
		IdleTimeout: time.Duration(opt.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", opt.Address, redis.DialDatabase(opt.DB))
			if err != nil {
				return nil, err
			}
			if opt.Password != "" {
				if _, err := c.Do("AUTH", opt.Password); err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

/*** Single KV invoke ***/
func (p *ConnPool) SetString(key string, value interface{}) (interface{}, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return conn.Do("SET", key, value)
}

func (p *ConnPool) GetString(key string) (string, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}

func (p *ConnPool) GetInt(key string) (int, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return redis.Int(conn.Do("GET", key))
}

func (p *ConnPool) GetInt64(key string) (int64, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("GET", key))
}

func (p *ConnPool) DelKey(key string) (interface{}, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return conn.Do("DEL", key)
}

func (p *ConnPool) ExpireKey(key string, expireTime int64) (interface{}, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return conn.Do("EXPIRE", key, expireTime)
}

func (p *ConnPool) HSet(key string, HKey string, data interface{}) (interface{}, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return conn.Do(REDIS_KEYWORD_HSET, key, HKey, data)
}

func (p *ConnPool) HGet(key string, HKey string) (interface{}, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return conn.Do(REDIS_KEYWORD_HGET, key, HKey)
}

func (p *ConnPool) HMGet(key string, hashKeys ...string) ([]interface{}, error) {
	conn := p.raw.Get()
	defer conn.Close()
	ret, err := conn.Do(REDIS_KEYWORD_HMGET, key, hashKeys)
	if err != nil {
		return nil, err
	}
	reta, ok := ret.([]interface{})
	if !ok {
		return nil, errors.New("result not an array")
	}
	return reta, nil
}

func (p *ConnPool) HMSet(key string, hashKeys []string, vals []interface{}) (interface{}, error) {
	if len(hashKeys) == 0 || len(hashKeys) != len(vals) {
		var ret interface{}
		return ret, errors.New("bad length")
	}
	input := []interface{}{key}
	for i, v := range hashKeys {
		input = append(input, v, vals[i])
	}
	conn := p.raw.Get()
	defer conn.Close()
	return conn.Do(REDIS_KEYWORD_HMSET, input...)
}

func (p *ConnPool) HGetString(key string, HKey string) (string, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return redis.String(conn.Do(REDIS_KEYWORD_HGET, key, HKey))
}
func (p *ConnPool) HGetFloat(key string, HKey string) (float64, error) {
	conn := p.raw.Get()
	defer conn.Close()
	f, err := redis.Float64(conn.Do(REDIS_KEYWORD_HGET, key, HKey))
	return float64(f), err
}
func (p *ConnPool) HGetInt(key string, HKey string) (int, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return redis.Int(conn.Do(REDIS_KEYWORD_HGET, key, HKey))
}
func (p *ConnPool) HGetInt64(key string, HKey string) (int64, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return redis.Int64(conn.Do(REDIS_KEYWORD_HGET, key, HKey))
}
func (p *ConnPool) HGetBool(key string, HKey string) (bool, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return redis.Bool(conn.Do(REDIS_KEYWORD_HGET, key, HKey))
}
func (p *ConnPool) HDel(key string, HKey string) (interface{}, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return conn.Do(REDIS_KEYWORD_HDEL, key, HKey)
}

func (p *ConnPool) HGetAll(key string) (interface{}, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return conn.Do(REDIS_KEYWORD_HGETALL, key)
}

func (p *ConnPool) HGetAllStringMap(key string) (map[string]string, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return redis.StringMap(conn.Do(REDIS_KEYWORD_HGETALL, key))
}

func (p *ConnPool) HGetAllValues(key string) ([]interface{}, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return redis.Values(conn.Do(REDIS_KEYWORD_HGETALL, key))
}
func (p *ConnPool) HGetAllString(key string) ([]string, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return redis.Strings(conn.Do(REDIS_KEYWORD_HGETALL, key))
}

func (p *ConnPool) Incr(key string) (int64, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return redis.Int64(conn.Do(REDIS_KEYWORD_INCR, key))
}

func (p *ConnPool) Decr(key string) (int64, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return redis.Int64(conn.Do(REDIS_KEYWORD_DECR, key))
}

func (p *ConnPool) IncrBy(key string, incBy int64) (int64, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return redis.Int64(conn.Do(REDIS_KEYWORD_INCRBY, key, incBy))
}

func (p *ConnPool) DecrBy(key string, decrBy int64) (int64, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return redis.Int64(conn.Do(REDIS_KEYWORD_DECRBY, key))
}

func (p *ConnPool) IncrByFloat(key string, incBy float64) (float64, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return redis.Float64(conn.Do(REDIS_KEYWORD_INCRBYFLOAT, key, incBy))
}

func (p *ConnPool) DecrByFloat(key string, decrBy float64) (float64, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return redis.Float64(conn.Do(REDIS_KEYWORD_DECRBYFLOAT, key, decrBy))
}

func (p *ConnPool) HIncrBy(key string, filed interface{}, increment interface{}) (uint64, error) {
	conn := p.raw.Get()
	defer conn.Close()
	return redis.Uint64(conn.Do(REDIS_KEYWORD_HINCRBY, key, filed, increment))
}

func (p *ConnPool) Scan(cursor int64, pattern string, count int64) (int64, []string, error) {
	conn := p.raw.Get()
	defer conn.Close()
	var items []string
	var newCursor int64

	values, err := redis.Values(conn.Do("SCAN", cursor, "MATCH", pattern, "COUNT", count))
	if err != nil {
		return 0, nil, err
	}
	values, err = redis.Scan(values, &newCursor, &items)
	if err != nil {
		return 0, nil, err
	}

	return newCursor, items, nil
}

func (p *ConnPool) HMGetStringMap(key string, fields ...string) (map[string]string, error) {
	conn := p.raw.Get()
	defer conn.Close()

	l := len(fields)
	cmd := append([]string{key}, fields...)

	payload := make([]interface{}, len(cmd))
	for n, c := range cmd {
		payload[n] = c
	}

	result := make(map[string]string)
	resp, err := redis.Strings(conn.Do(REDIS_KEYWORD_HMGET, payload...))
	if err != nil {
		return nil, err
	}

	var ctrFail int
	for i := 0; i < l; i++ {
		result[fields[i]] = resp[i]
		if resp[i] == "" {
			ctrFail++
		}
	}

	if ctrFail == l {
		return result, NotFoundErr
	}

	return result, nil
}

func (p *ConnPool) GetConnection() (redis.Conn, error) {
	conn := p.raw.Get()
	return conn, conn.Err()
}
