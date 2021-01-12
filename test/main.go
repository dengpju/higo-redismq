package main

import (
	"fmt"
	"github.com/dengpju/higo-redis/redis"
	"github.com/dengpju/higo-utils/utils"
	"math/rand"
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
			redis.PoolMaxIdle(11),
			redis.PoolMaxIdleTime(60),
			redis.PoolMaxConnLifetime(10),
			redis.PoolWait(true),
		))
	/**
	c := make(chan int)
	go producer(c)
	r := consumer(c)
	<-r // 输出结果，会在此卡住

	 */
	go producer1()
	go consumer1()
	go consumer1()
	go consumer1()
	select {

	}
}

func producer(out chan int)  {
	defer close(out)
	for  {
		i := rand.Intn(1000) + 1
		redis.Redis.Lpush("orders", i)
		fmt.Printf("生产订单：%d\n",i)
		out <- i
		time.Sleep(time.Second * 3)
	}
}

func consumer(out chan int) (r chan struct{}) {
	r = make(chan struct{})
	go func() {
		defer func() {
			r <- struct{}{}
		}()
		for item := range out{// 会阻塞，卡住等待
			fmt.Printf("消费订单：%d\n",item)
			//order := redis.Redis.Rpop("orders")
			//fmt.Printf("redis订单：%s\n",order)
		}
	}()

	return r
}

func producer1()  {
	for  {
		go func() {
			redis.New(
				redis.NewPoolConfigure(
					redis.PoolHost("192.168.8.99"),
					redis.PoolPort(6379),
					redis.PoolAuth("1qaz2wsx"),
					redis.PoolDb(0),
					redis.PoolMaxConnections(100),
					redis.PoolMaxIdle(11),
					redis.PoolMaxIdleTime(60),
					redis.PoolMaxConnLifetime(10),
					redis.PoolWait(true),
				))
			gid := utils.GoroutineID()
			i := rand.Intn(1000) + 1
			redis.Redis.Lpush("orders", i)
			fmt.Printf("1==生产者 协程id：%d 生产订单：%d\n",gid, i)
			time.Sleep(time.Second * 1)
		}()
	}

}

func consumer1() {
	for  {
		go func() {
			redis.New(
				redis.NewPoolConfigure(
					redis.PoolHost("192.168.8.99"),
					redis.PoolPort(6379),
					redis.PoolAuth("1qaz2wsx"),
					redis.PoolDb(0),
					redis.PoolMaxConnections(100),
					redis.PoolMaxIdle(11),
					redis.PoolMaxIdleTime(60),
					redis.PoolMaxConnLifetime(10),
					redis.PoolWait(true),
				))
			gid := utils.GoroutineID()
			order := redis.Redis.Rpop("orders")
			fmt.Printf("2==消费者 协程id：%d 消费订单：%s\n",gid, order)
			time.Sleep(time.Second * 1)
		}()
	}
}