package main

import (
	"fmt"   
	//"log"   
	"strconv"
	"sync"
	"time"
)

var fmeng = false

func producer(threadID int, wg *sync.WaitGroup, ch chan string)  {
	count := 0
	for !fmeng {
		time.Sleep(time.Second * 1)
		count ++
		data := strconv.Itoa(threadID) + "---" + strconv.Itoa(count)
		fmt.Printf("producer, %s\n", data)
		ch <- data
	}
	wg.Done()
}

func consumer(wg *sync.WaitGroup, ch chan string)  {
	for data := range ch {
		time.Sleep(time.Second * 1)
		fmt.Printf("consumer, %s\n", data)
	}
	wg.Done()
}


func main() {
	chanSteam := make(chan string, 10)

	wgPd := new(sync.WaitGroup)
	wgCs := new(sync.WaitGroup)

	for i := 0; i < 3; i ++ {
		wgPd.Add(1)
		go producer(i, wgPd, chanSteam)
	}

	for j := 0; j < 2; j ++ {
		wgCs.Add(1)
		go consumer(wgCs, chanSteam)
	}

	go func() {
		time.Sleep(time.Second * 3)
		fmeng = true
	}()

	wgPd.Wait()

	close(chanSteam)
	wgCs.Wait()

}
