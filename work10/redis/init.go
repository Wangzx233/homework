package redisInit

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	Addr        = "localhost:6379"
	IdLeTimeout = 5
	MaxIdle     = 20
	MaxActive   = 8
)

type OptionPool struct {
	addr        string
	idLeTimeout int
	maxIdle     int
	maxActive   int
}

type PoolExt interface {
	apply(pool *OptionPool)
}

type tempFunc func(pool *OptionPool)

type funcPoolExt struct {
	t tempFunc
}

func (f funcPoolExt) apply(pool *OptionPool) {
	f.t(pool)
}

func NewFuncPoolExt(t tempFunc) *funcPoolExt {
	return &funcPoolExt{t}
}

type Client struct {
	Option OptionPool
	pool   *redis.Pool
}

var DefaultOption = OptionPool{
	addr:        Addr,
	idLeTimeout: IdLeTimeout,
	maxIdle:     MaxIdle,
	maxActive:   MaxActive,
}

func NewClient(op ...PoolExt) *Client {
	c := &Client{Option: DefaultOption}
	for _, p := range op {
		p.apply(&c.Option)
	}
	c.setRedisPool()
	return c
}

func (pc *Client) setRedisPool() {
	pc.pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", pc.Option.addr)
			if conn == nil || err != nil {
				return nil, err
			}
			return conn, nil
		},
		MaxIdle:     pc.Option.maxIdle,                                  // 最大空闲连接数
		MaxActive:   pc.Option.maxActive,                                // 最大活跃连接数
		IdleTimeout: time.Second * time.Duration(pc.Option.idLeTimeout), // 连接等待时间
	}
}

//字符串的set和get
func (pc *Client) Set(args ...interface{}) error {
	c := pc.pool.Get()
	defer c.Close()
	_, err := c.Do("SET", args...)
	if err != nil {
		return err
	}
	return err
}

func (pc *Client) Get(args ...interface{}) (interface{}, error) {
	c := pc.pool.Get()
	defer c.Close()
	value, err := c.Do("GET", args...)
	if err != nil {
		return value, err
	}
	return value, err
}

//哈希的get和set
func (pc *Client) HSet(args ...interface{}) error {
	c := pc.pool.Get()
	defer c.Close()
	_, err := c.Do("HSET", args...)
	if err != nil {
		return err
	}
	return err
}
func (pc *Client) HGet(args ...interface{}) (interface{}, error) {
	c := pc.pool.Get()
	defer c.Close()
	value, err := c.Do("HGET", args...)
	if err != nil {
		return value, err
	}
	return value, err
}


//哈希的get和set
func (pc *Client) LPUSH(args ...interface{}) error {
	c := pc.pool.Get()
	defer c.Close()
	_, err := c.Do("LPUSH", args...)
	if err != nil {
		return err
	}
	return err
}
func (pc *Client) LPOP(args ...interface{}) (interface{}, error) {
	c := pc.pool.Get()
	defer c.Close()
	value, err := c.Do("LPOP", args...)
	if err != nil {
		return value, err
	}
	return value, err
}
func (pc *Client) LRANGE(args ...interface{}) (interface{}, error) {
	c := pc.pool.Get()
	defer c.Close()
	value, err := c.Do("LRANGE", args...)
	if err != nil {
		return value, err
	}
	return value, err
}
