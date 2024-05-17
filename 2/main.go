package main

import (
	"fmt"
	"time"
)

/*
func contador(x int)  {
	for i := 0; i < x; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() { 		//T1
	go contador(10)		//T2
	go contador(10)		//T3
	contador(10)	//T1
}

func main() { //T1
	canal := make(chan string)

	go func() { //T2
		canal <- "opa"
	} ()

	msg := <- canal
	fmt.Println(msg) //T1
}
*/

func worker(workerId int, data chan int){
	for x := range data { // Ler do canal
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main()  { //T1
	canal := make(chan int)

	qtdWorker := 10

	for i := 0; i < qtdWorker; i++ {
		go worker(i, canal) //Ti
	}

	/*
	go worker(1, canal) //T2
	go worker(2, canal) //T3
	go worker(3, canal) //T4
	go worker(4, canal) //T5
	*/

	for i := 0; i < 10000; i++ {
		canal <- i
	}
}