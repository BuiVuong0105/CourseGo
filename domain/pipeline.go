package domain

import "fmt"

type Pipeline struct{}

func CaculateSQRT() {
	pipeline := &Pipeline{}
	pipeline.sumQRT()
}

func (p *Pipeline) sumQRT() {
	randomNumbers := []int{}
	for i := 1; i < 5; i++ {
		randomNumbers = append(randomNumbers, i)
	}

	// Gen Pipeline
	pipelineChannel := p.generatePipeline(randomNumbers)
	fanoutFirst := p.fanout(pipelineChannel)
	fanoutSecond := p.fanout(pipelineChannel)
	fanoutThird := p.fanout(pipelineChannel)

	fanIn := p.fanIn(fanoutFirst, fanoutSecond, fanoutThird)

	sum := 0

	for res := range fanIn {
		sum += res
	}

	fmt.Println("Resule: ", sum)
}

func (p *Pipeline) generatePipeline(numbers []int) <-chan int {
	channel := make(chan int)
	go func() {
		for _, number := range numbers {
			channel <- number
		}
		close(channel)
	}()
	return channel
}

func (p *Pipeline) fanout(channelIn <-chan int) <-chan int {
	channel := make(chan int)
	go func() {
		for res := range channelIn {
			channel <- res * res
		}
		close(channel)
	}()
	return channel
}

func (p *Pipeline) fanIn(channelIn ...<-chan int) <-chan int {
	channel := make(chan int)
	go func() {
		for _, channelDetail := range channelIn {
			for res := range channelDetail {
				channel <- res * res
			}
		}
		close(channel)
	}()
	return channel
}
