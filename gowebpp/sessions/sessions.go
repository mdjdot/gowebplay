package sessions

import (
	"gowebpp/confs"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

// Sessions 会话列表
// var Sessions = make(map[string]*time.Timer)
// var locker = sync.Mutex{}

// func init() {
// 	go func() {
// 		for {
// 			time.Sleep(5*time.Second)
// 			for token, timer := range Sessions {
// 				go func(to string, t *time.Timer) {
// 					<-t.C
// 					locker.Lock()
// 					delete(Sessions, to)
// 					locker.Unlock()

// 				}(token, timer)
// 			}
// 		}
// 	}()
// }

// Add 添加 session
func Add(token string, ex int) error {
	conn := confs.RedisPool.Get()
	_, err := conn.Do("AUTH", "dmtest")
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", token, "1", "EX", strconv.Itoa(ex))
	if err != nil {
		return err
	}
	return nil
}

// Get 获取 session
func Get(token string) (string, error) {
	conn := confs.RedisPool.Get()
	_, err := conn.Do("AUTH", "dmtest")
	if err != nil {
		return "", err
	}

	value, err := redis.String(conn.Do("GET", token))
	if err != nil {
		return "", err
	}
	return value, err
}
