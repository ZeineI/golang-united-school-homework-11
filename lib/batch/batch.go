package batch

import (
	"fmt"
	"sync"
)

type user struct {
	ID int64
}

func getOne(wg *sync.WaitGroup, poolNum int64, receive <-chan int64, response chan<- user) {
	defer wg.Done()
	for id := range receive {
		fmt.Printf("worker-%d take element %d\n", poolNum, id)
		response <- user{ID: id}
	}
}

func getBatch(n int64, pool int64) (res []user) { // n - num of users //pool - num of goroutines

	wg := &sync.WaitGroup{}
	wg.Add(int(pool))

	//channel
	receive := make(chan int64)
	response := make(chan user)

	for i := int64(0); i < pool; i++ {
		go getOne(wg, i, receive, response)
	}

	for i := int64(0); i < n; i++ {
		receive <- i
		res = append(res, <-response)
	}

	//done
	close(receive)
	close(response)
	wg.Wait()
	return res
}
