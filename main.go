package main

import (
	//	"time"
	//	"DynamicKey5"
	"fmt"
	"runtime"
	"sync"
	"time"
	"github.com/sajari/word2vec"
	"log"
)

var (
	count int32
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	c := word2vec.Client{Addr: "localhost:1234"}

	// Create an expression.
	expr := word2vec.Expr{}
	expr.Add(1, "PC")
	//expr.Add(1, "crap")
	//expr.Add(1, "woman")

	// Find the most similar result by cosine similarity.
	matches, err := c.CosN(expr, 10)
	if err != nil {
		log.Fatalf("error evaluating cosine similarity: %v", err)
	}else {
		for i,v:=range matches{
			fmt.Println(i,v)
		}
	}
}
func mut() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
}
func incCount() {
	defer wg.Done()
	for i := 0; i < 4; i++ {
		mutex.Lock()
		value := count
		runtime.Gosched()
		value++
		count = value
		mutex.Unlock()
	}
}

func chann() {
	ch := make(chan int)
	go func() {
		var sum int = 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		ch <- sum
		time.Sleep(time.Second)
		fmt.Println("Im before the <-ch")
	}()

	fmt.Println(<-ch)
	fmt.Println("Im after the <-ch")
	time.Sleep(time.Second*2)
}

//two rely on the channel one's value,so wait until channel one finish the execution,even the one channel start after one second sleep
//channel could auto wait the other channel that need to be finished
//if the channel has been <- channelName ,then you can't use select or execute <- channelName again
func channw() {
	one := make(chan int)
	two := make(chan int)
	three := make(chan int)
	go func() {
		v := <-one //这会导致直接等待one执行后才会调用select ，导致 three 也执行完了 ，然后 输出的结果就是1000
		two <- v   //只要two一旦被赋值 就会立马执行最后面的输出语句
		select {
		case x0 := <-one:
			two <- x0
		case x1 := <-three://会走到这里是因为three 还没有被取出过值，而one前面已经被取出来
			two <- x1
			two<-10000
		}

	}()
	go func() {
		time.Sleep(time.Second * 5)
		one <- 5000
	}()
	go func() {
		time.Sleep(time.Second * 2)
		three <- 2000
	}()
	fmt.Println("got two", <-two) // 输出5000，是第一个channel 赋值传过来的数值
	fmt.Println("got two", <-two) //输出2000，是第二个select执行第二地赋值传过来的数值
	//fmt.Println ("got two",<-two)//err,all groutines are asleep,因为two的channel只被赋值过两次，第二次已经被取出来了
	//如果上面的two再进行一次channel的赋值 ，则上一条语句不会出错
}