package uuid

import (
	"encoding/hex"
	idworker "github.com/gitstliu/go-id-worker"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var GenerateRandMap map[string]string
var GenerateRandRWMutex *sync.RWMutex

func NewId() int {
	worker := &idworker.IdWorker{}
	err := worker.InitIdWorker(1000, 10000)
	if err != nil {
		log.Print("NewId Error", err)
	}

	newId, newIdErr := worker.NextId()
	log.Println(newId, newIdErr)
	return int(newId)
}

type AutoInc struct {
	start, step int
	queue       chan int
	running     bool
}

func NewAutoIncInit(start, step int) (ai *AutoInc) {
	ai = &AutoInc{
		start:   start,
		step:    step,
		running: true,
		queue:   make(chan int, 4),
	}
	go ai.process()
	return
}

func (ai *AutoInc) process() {
	defer func() { recover() }()
	for i := ai.start; ai.running; i = i + ai.step {
		ai.queue <- i
	}
}

func (ai *AutoInc) Id() int {
	return <-ai.queue
}

func (ai *AutoInc) Close() {
	ai.running = false
	close(ai.queue)
}

var AutoId *AutoInc

// 20211102195043100000
func NewAutoIdInc(aid int, intprefix string) string {
	b := strconv.Itoa(aid)
	//log.Println(DateTimeNowFormat()+b, reflect.TypeOf(b))
	return intprefix + time.Now().Format("20060102150405") + b
}

var numberletters = []byte("123456789")
var letters = []byte("abcdefghjkmnpqrstuvwxyz123456789")
var longLetters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().Unix())
}

// RandLow 随机字符串，包含 1~9 和 a~z - [i,l,o]
func RandLow(n int) []byte {
	if n <= 0 {
		return []byte{}
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return []byte{}
	}
	for i, x := range b {
		arc = x % 31
		b[i] = letters[arc]
	}
	return b
}

// RandLow 随机字符串，包含 1~9
func RandNumber(n int) []byte {
	if n <= 0 {
		return []byte{}
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return []byte{}
	}
	for i, x := range b {
		arc = x % 9
		b[i] = numberletters[arc]
	}
	return b
}

// RandUp 随机字符串，包含 英文字母和数字附加=_两个符号
func RandUp(n int) []byte {
	if n <= 0 {
		return []byte{}
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return []byte{}
	}
	for i, x := range b {
		arc = x % 61
		b[i] = longLetters[arc]
	}
	return b
}

// RandHex 生成16进制格式的随机字符串
func RandHex(n int) []byte {
	if n <= 0 {
		return []byte{}
	}
	var need int
	if n&1 == 0 { // even
		need = n
	} else { // odd
		need = n + 1
	}
	size := need / 2
	dst := make([]byte, need)
	src := dst[size:]
	if _, err := rand.Read(src[:]); err != nil {
		return []byte{}
	}
	hex.Encode(dst, src)
	return dst[:n]
}

func Rands(namber int) []string {
	rands := make([]string, 0)
	for i := 0; i < namber; i++ {
		rands = append(rands, string(RandLow(8)))
	}
	return rands
}

func NumberRands(namber int) []string {
	rands := make([]string, 0)
	for i := 0; i < namber; i++ {
		rands = append(rands, string(RandNumber(8)))
	}
	return rands
}

// 数据同步的请求ID
func GenerateRand(key string) string {
	GenerateRandRWMutex.Lock()
GenerateRand:
	Rand := string(RandUp(10))
	str := Rand + key
	value, ok := GenerateRandMap[str]
	if ok {
		log.Errorln("重复ID value:", value)
		goto GenerateRand
	}
	GenerateRandMap[str] = str
	log.Println("len(GenerateRandMap):", len(GenerateRandMap), "Rand+key:", str)
	GenerateRandRWMutex.Unlock()
	return str
}
