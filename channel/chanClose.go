package chanClose

import (
	"fmt"
	"time"
)

func main() {

	jobs := make(chan int, 4)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received Jobs ", j)
				time.Sleep(1 * time.Second)
			} else {
				fmt.Println("All jobs received....")
				done <- true
				return
			}
		}
	}()
	for k := 0; k <= 4; k++ {
		jobs <- k
		fmt.Println("JobSending .. ", k)
	}
	close(jobs)
	fmt.Println("All Jobs Send.. ")
	<-done

	q := make(chan int, 4)

	q <- 1
	q <- 2
	close(q)

	for elem := range q {
		fmt.Println(elem)
	}

}
