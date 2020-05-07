package sessions

import (
	"sync"
	"time"
)

// Sessions 会话列表
var Sessions = make(map[string]*time.Timer)
var locker = sync.Mutex{}

func init() {
	go func() {
		for {
			time.Sleep(5*time.Second)
			for token, timer := range Sessions {
				go func(to string, t *time.Timer) {
					<-t.C
					locker.Lock()
					delete(Sessions, to)
					locker.Unlock()

				}(token, timer)
			}
		}
	}()
}
