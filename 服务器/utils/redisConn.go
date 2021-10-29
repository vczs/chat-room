package utils

import (
	"github.com/garyburd/redigo/redis"
)

var(
	Pool *redis.Pool
)

func InitPool(address string , maxIdle int , maxActive int) {
	//创建redis连接池并放入redis链接
	Pool = &redis.Pool{
		MaxIdle : maxIdle ,                                                //最大空闲链接数
		MaxActive : maxActive ,                                            //和redis数据库的最大连接数 0表示没有限制
		IdleTimeout : 100 ,                                                //链接的最大空闲时间 超过100s链接未使用就放入连接池中
		Dial: func()(redis.Conn,error){return redis.Dial("tcp",address)} , //创建redis连接池
	}
}
