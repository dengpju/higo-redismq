package main

import (
	"fmt"
	"github.com/dengpju/higo-redis/redis"
	"github.com/dengpju/higo-utils/utils"
	"time"
)

func main()  {
	redis.New(
		redis.NewPoolConfigure(
			redis.PoolHost("192.168.8.99"),
			redis.PoolPort(6379),
			redis.PoolAuth("1qaz2wsx"),
			redis.PoolDb(0),
			redis.PoolMaxConnections(100),
			redis.PoolMaxIdle(3),
			redis.PoolMaxIdleTime(60),
			redis.PoolMaxConnLifetime(10),
			redis.PoolWait(true),
		))
	go consumer1()
	select {

	}
}

func consumer1() {
	for  {
		gid := utils.GoroutineID()
		order := redis.Redis.Rpop("orders")
		fmt.Printf("2==消费者 协程id：%d 消费订单：%s\n",gid, order)
		time.Sleep(time.Microsecond * 500)
	}
}