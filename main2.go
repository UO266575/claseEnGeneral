package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Starting the application...")
	c1 := make(chan int)
	c2 := make(chan int)
	defer close(c1)
	defer close(c2)

	go firstNum(c1, &wg)
	go secondNum(c2, &wg)

	counter := 0
	for counter < 20 {
		wg.Add(1)
		select {
		case msg1 := <-c1:
			if msg1%2 == 0 {
				go firstReceiver(c1, &wg)
			} else {
				go secondReceiver(c1, &wg)
			}
		case msg2 := <-c2:
			if msg2%2 == 0 {
				go firstReceiver(c2, &wg)
			} else {
				go secondReceiver(c2, &wg)
			}
		}
		counter++
		wg.Wait()
	}
	fmt.Println("Receiver 1  -  ", nR1)
	fmt.Println("Receiver 2  -  ", nR2)
}

var nR1 []int

func firstReceiver(ch <-chan int, wg *sync.WaitGroup) {
	val := <-ch
	nR1 = append(nR1, val)
	//fmt.Println("Receiver 1  -  ", nR1)
	wg.Done()
}

var nR2 []int

func secondReceiver(ch2 <-chan int, wg *sync.WaitGroup) {
	val2 := <-ch2
	nR2 = append(nR2, val2)
	//fmt.Println("Receiver 2  -  ", nR2)
	wg.Done()
}

func firstNum(c1 chan<- int, wg *sync.WaitGroup) {
	for i := 0; true; i++ {
		c1 <- i
		time.Sleep(time.Millisecond * 1000)
	}
	wg.Done()
	close(c1)
}

func secondNum(c2 chan<- int, wg *sync.WaitGroup) {
	for j := 11; true; j++ {
		c2 <- j
		time.Sleep(time.Millisecond * 500)
	}
	wg.Done()
	close(c2)
}

// go firstNum(c1)
// msg := <-c1
// fmt.Println("El num es: ", msg)
// if msg%2 == 0 {
// 	fmt.Println("PAR")
// 	go firstReceiver(c1)
// 	counter++
// 	if counter >= 5 {
// 		break
// 	}
// }
// } else {
// 	fmt.Println("IMPAR")
// 	go secondReceiver(c1)
// 	if counter >= 5 {
// 		break
// 	}
// }

// go secondNum(c2)
// msg2 := <-c2
// fmt.Println("El num es: ", msg2)
// if msg2%2 == 0 {
// 	fmt.Println("PAR")
// 	go firstReceiver(c2)
// 	if counter >= 5 {
// 		break
// 	}
// } else {
// 	fmt.Println("IMPAR")
// 	go secondReceiver(c2)
// 	if counter >= 5 {
// 		break
// 	}
// }
// }
