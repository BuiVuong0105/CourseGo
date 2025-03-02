package concurrency

import (
	"log"
	"sync"
	"time"
)

func HeavyTask(weight int32) {
	time.Sleep(time.Second * time.Duration(weight))
}

// Nhiều luồng chạy cùng truy cập 1 biến dẫn đến RaceCondition
func raceConditionAndDataRace() {
	var counter int32
	go func() {
		counter++ // truy cập thằng biến counter đã khai báo, không phải bản copy
	}()
	HeavyTask(1)
	log.Printf("the value is %v \n", counter)
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>MUTEX. Thư viện trong gói syn giúp loại trừ truy cập từ nhiều grountine, giúp truy cập thẳng share mem chứ ko truy cập vào cpu cache
// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex // Là object thì mặc định đã được khởi tạo
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func runMutex() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	HeavyTask(5)
	log.Printf("Value Of Somkey: %v", c.Value("somekey"))
}

// END MUTEX >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

// START WAIT GROUP >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// Tương tự như countdownlatch ở Java, cho phép 1 nhóm luồng có thể  truy cập 1 đoạn code quan trọng
func RunWaitGroup() {
	wg := sync.WaitGroup{}
	for index, element := range []string{"hello", "greetings", "good day"} {
		log.Printf("Index: %v, Element: %v", index, element)
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			log.Printf("Salution: %s", salutation)
		}(element)
	}
	wg.Wait()
}

// END WAIT GROUP >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func RunCommonConcurrency() {
	// raceConditionAndDataRace()
	// runMutex()
	RunWaitGroup()
}
