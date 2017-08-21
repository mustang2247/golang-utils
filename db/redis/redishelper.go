package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

var client *redis.Client

func Open(host string, pwd string, dbNum int) (error) {
	client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pwd, // no password set
		DB:       dbNum,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return err
}

func Set(key string, value string) (error) {
	err := client.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}

	return err
}

func Get(key string) (string, error) {
	val, err := client.Get(key).Result()
	if err == redis.Nil {
		fmt.Println("key does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key", val)
	}

	return val, err
}

/*检查错误信息*/
func checkErr(err error) {
	if err != nil {
		println("mysql checkErr: " + err.Error())
		panic(err)
	}
}