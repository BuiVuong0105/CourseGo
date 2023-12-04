package singleton

import (
	"fmt"
	"sync"
	"time"
)

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

var (
	instance *singleton
	lock     sync.Once
)

func init() {
	instance = &singleton{count: 100}
	fmt.Printf("INIT SINGLETON: %p \n", instance)
}

func GetInstance() Singleton {

	if instance != nil {
		return instance
	}

	lock.Do(func() {
		if instance == nil {
			time.Sleep(500 * time.Millisecond)
			instance = &singleton{count: 100}
		}
	})

	return instance
}

func TestSingletonPattern() {
	for i := 1; i <= 10; i++ {
		go func() {
			singleton := GetInstance()
			fmt.Printf("%p \n", singleton)

		}()
	}
	time.Sleep(5 * time.Second)
}
