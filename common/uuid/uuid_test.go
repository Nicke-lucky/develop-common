package uuid

import (
	"log"
	"sync"
	"testing"
	"time"
)

var (
	UserIdMap map[string]string

	UserIdMapRWMutex sync.RWMutex
)

func TestGetUuid(t *testing.T) {
	uuid := GetUuid("")
	log.Println("uuid:", uuid, len(uuid))
	//45位
	uuid = GetUuid("3301" + time.Now().Format("20060102150405"))
	log.Println("uuid:", uuid, len(uuid))
	UserIdMap = make(map[string]string, 100)

}

// fatal error: concurrent map read and map write
func TestConcurrentMap(t *testing.T) {
	uuid := GetUuid("")
	log.Println("uuid:", uuid, len(uuid))
	//45位
	uuid = GetUuid("3301" + time.Now().Format("20060102150405"))
	log.Println("uuid:", uuid, len(uuid))
	UserIdMap = make(map[string]string, 100)
	go GetMap()

	go ReadMapIs()
	go ReadMap()
	for {
		time.Sleep(time.Second * 1)
		if len(UserIdMap) > 500 {
			log.Println("加锁后，map并发读写问题解决")
			return
		}
	}
}
func TestConcurrentMapRWMutex(t *testing.T) {
	uuid := GetUuid("")
	log.Println("uuid:", uuid, len(uuid))
	//45位
	uuid = GetUuid("3301" + time.Now().Format("20060102150405"))
	log.Println("uuid:", uuid, len(uuid))
	UserIdMap = make(map[string]string, 100)
	//map并发读写,解决办法：加读写锁
	go GetMapRWMutex()
	go ReadMapIsRWMutex()
	go ReadMapRWMutex()
	for {
		time.Sleep(time.Second * 1)
		if len(UserIdMap) > 500 {
			log.Println("加锁后，map并发读写问题解决")
			return
		}
	}
}
func GetMap() {
	for {
		uuid := GetUuid("")
		//fatal error: concurrent map iteration and map write
		UserIdMap[uuid] = uuid
		time.Sleep(time.Microsecond * 100)
	}
}

func ReadMapIs() {
	for {
		v, ok := UserIdMap["key"]
		if ok {
			log.Println("ok,v:", v)
		} else {
			log.Println("!ok")
		}
		time.Sleep(time.Microsecond * 100)
	}
}

func ReadMap() {
	for {
		//fatal error: concurrent map iteration and map write
		for k, v := range UserIdMap {
			log.Println("k,v:", k, v)
		}
		time.Sleep(time.Microsecond * 100)
	}
}

func GetMapRWMutex() {
	for {
		uuid := GetUuid("")
		//fatal error: concurrent map iteration and map write
		//解决办法：加读写锁

		UserIdMapRWMutex.Lock()
		UserIdMap[uuid] = uuid
		UserIdMapRWMutex.Unlock()
		time.Sleep(time.Microsecond * 100)
	}
}

func ReadMapIsRWMutex() {
	for {
		UserIdMapRWMutex.RLock()
		v, ok := UserIdMap["key"]
		UserIdMapRWMutex.RUnlock()
		if ok {
			log.Println("ok,v:", v)
		} else {
			log.Println("!ok")
		}
		time.Sleep(time.Microsecond * 100)
	}
}

func ReadMapRWMutex() {
	for {
		//fatal error: concurrent map iteration and map write

		UserIdMapRWMutex.RLock()
		for k, v := range UserIdMap {
			log.Println("k,v:", k, v)
		}
		UserIdMapRWMutex.RUnlock()
		time.Sleep(time.Microsecond * 100)
	}
}
