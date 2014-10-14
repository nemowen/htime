package dao

import (
	redis "github.com/alphazero/Go-Redis" //这个比较好用
)

type RedisDao struct {
	client redis.AsyncClient
}
