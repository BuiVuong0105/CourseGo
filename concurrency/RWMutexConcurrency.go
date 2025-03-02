package concurrency

import (
	"fmt"
	"log"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/google/uuid"
)

type ShareCounter struct {
	counter int32
}

type RShare ShareCounter

func valueOf(shareCounter *ShareCounter, transId string, lock sync.Locker, waitGroup *sync.WaitGroup) {
	lock.Lock()
	defer lock.Unlock()
	defer waitGroup.Done()
	log.Printf("[%v] Read Value Of ShareCounter: %v", transId, shareCounter.counter)
}

func increment(shareCounter *ShareCounter, number int32, transId string, lock sync.Locker, waitGroup *sync.WaitGroup) {
	lock.Lock()
	defer lock.Unlock()
	defer waitGroup.Done()
	shareCounter.counter = shareCounter.counter + number
	log.Printf("[%v] Increment Success: %v", transId, number)
}

func RunRWMutextV2() {
	waitGroup := &sync.WaitGroup{}
	lock := sync.RWMutex{}
	shareCounter := ShareCounter{}
	transIds := make([]string, 0, 10)
	transIds = append(transIds, uuid.New().String())
	transIds = append(transIds, uuid.New().String())
	transIds = append(transIds, uuid.New().String())
	transIds = append(transIds, uuid.New().String())
	transIds = append(transIds, uuid.New().String())
	transIds = append(transIds, uuid.New().String())
	transIds = append(transIds, uuid.New().String())

	for index, transId := range transIds {
		waitGroup.Add(1)
		go increment(&shareCounter, int32(index), transId, &lock, waitGroup)
	}

	for _, transId := range transIds {
		waitGroup.Add(1)
		go valueOf(&shareCounter, transId, lock.RLocker(), waitGroup)
	}
	waitGroup.Wait()
}

func RunRWMutext() {
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1)
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}

		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")

	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}

}
