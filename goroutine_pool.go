package main

import (
	"fmt"
	"sync"
	"time"
)

func goPool() {
	start := time.Now()
	wg := new(sync.WaitGroup)
	data := make(chan int, 100)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			for v := range data {
				fmt.Println("goroutine:", n, v)
			}
		}(i)
	}

	for i := 0; i < 10000; i++ {
		data <- i
	}

	close(data)
	wg.Wait()		//阻塞主线程执行，直到所有的WaitGroup数量变成0
	end := time.Now()
	fmt.Println("goPool run time: ", end.Sub(start))
}

func noPool(){
	start := time.Now()
	wg := new(sync.WaitGroup)
	for i:=0;i<10000 ;i++  {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			println("nopool:", n)
		}(i)
	}
	wg.Wait()

	end := time.Now()
	fmt.Println("noPool run time: ",end.Sub(start))
}

type worker struct{
	working func(i int)
}

/**

goroutine池里面goroutine数目和核数有没有关系

1.goroutine池跑不满。
意味着channel data一有数据就会被goroutine拿走，这样的话当然只要CPU能调度的过来就行，也就是池子里的goroutine数目和CPU核数是最优的。经测试，确实是这样。

2.channel data有数据阻塞。
就是说goroutine是不够用的，如果goroutine的运行任务不是CPU密集型的（大部分情况都不是），而只是IO阻塞，这个时候一般goroutine数目在一定范围内是越多越好，当然范围在什么地方就要具体情况具体分析了。

*/
func closureFuncPool(){

	start := time.Now()

	wg:= new(sync.WaitGroup)

	changel := make(chan worker, 10)

	for i:=0;i<5 ;i++  {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ch := range changel{
				ch.working(1)
			}

		}()
	}

	for i:=0;i<10000 ;i++  {
		j := i
		w := worker{
			working: func(i int) {
				println(j,i)
			},
		}
		changel <- w
	}

	close(changel)
	wg.Wait()

	fmt.Println("closureFuncPool run time：",time.Now().Sub(start))
}

func main()  {

	//goPool()

	//noPool()

	closureFuncPool()
}