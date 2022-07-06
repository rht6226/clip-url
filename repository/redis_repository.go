package repository

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	apperrors "github.com/rht6226/clip-url/errors"
	"github.com/rht6226/clip-url/model"
)

type redisRepository struct {
	pool *redis.Pool
}

func NewRedisRepository(host, port string) (*redisRepository, error) {
	pool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
		},
	}

	return &redisRepository{pool: pool}, nil
}

// check if a uint is already being used by redis
func (r *redisRepository) IsUsed(id uint64) bool {
	conn := r.pool.Get()
	defer conn.Close()

	key := fmt.Sprintf("SHORTEN:%s", strconv.FormatUint(id, 10))
	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}

// save a new url entry to redis
func (r *redisRepository) Save(url *model.Url, t time.Time) error {
	conn := r.pool.Get()
	defer conn.Close()

	key := fmt.Sprintf("SHORTEN:%s", strconv.FormatUint(url.Id, 10))

	_, err := conn.Do("HMSET", redis.Args{key}.AddFlat(url))
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIREAT", key, t.Unix())

	if err != nil {
		return err
	}

	return nil
}

// load a url entry from the redis db
func (r *redisRepository) Load(id uint64) (*model.Url, error) {
	conn := r.pool.Get()
	defer conn.Close()

	key := fmt.Sprintf("SHORTEN:%s", strconv.FormatUint(id, 10))

	values, err := redis.Values(conn.Do("HGETALL", key))
	if err != nil {
		return nil, err
	} else if len(values) == 0 {
		return nil, apperrors.NewNotFound("URL does not exists")
	}

	var urlItem model.Url
	err = redis.ScanStruct(values, &urlItem)
	if err != nil {
		return nil, err
	}

	return &urlItem, nil
}

// close the connection pool
func (r *redisRepository) Close() error {
	return r.pool.Close()
}
