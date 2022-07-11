package batch

import (
	"fmt"
	"time"
)

type user struct {
	ID int64
}

func getOne( /*wg *sync.WaitGroup,*/ poolNum int64, receive <-chan int64, response chan<- user) {
	// defer wg.Done()
	for id := range receive {
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("Pool %d took user number %d\n", poolNum, id)
		response <- user{ID: id}
	}
}

func getBatch(n int64, pool int64) (res []user) { // n - num of users //pool - num of goroutines

	// wg := &sync.WaitGroup{}
	// wg.Add(int(pool))

	// channel
	receive := make(chan int64, pool)
	response := make(chan user, n)

	for i := int64(0); i < pool; i++ {
		go getOne( /*wg, */ i, receive, response)
	}

	for i := int64(0); i < n; i++ {
		receive <- i
	}

	// done

	close(receive)

	// wg.Wait()

	for i := int64(0); i < n; i++ {
		res = append(res, <-response)
	}
	close(response)

	return
}

// second implementation way

// func getOne(id int64) user {
// 	time.Sleep(time.Millisecond * 100)
// 	return user{ID: id}
// }

// func getBatch(n int64, pool int64) (res []user) {
// 	mx := &sync.Mutex{}

// 	errG, _ := errgroup.WithContext(context.Background())
// 	errG.SetLimit(int(pool))
// 	for i := int64(0); i < 100; i++ {
// 		iN := i
// 		errG.Go(func() error {
// 			resOne := getOne(iN)
// 			mx.Lock()
// 			defer mx.Unlock()
// 			res = append(res, resOne)

// 			return nil
// 		})
// 	}
// 	_ = errG.Wait()

// 	return res
// }

// func main() {
// 	fmt.Println(getBatch(100, 10))
// }
