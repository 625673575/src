package main

import (
	//	"time"
	//	"DynamicKey5"
	"fmt"
	"runtime"
	"sync"
	"time"
)

//func main() {
//
//	appID:="da54f706a6ea4f3c900018bacd48282f"
//	appCertificate:="0d14d6bc92984055b8d281e185894676"
//	channelName := "liu"
//	unixTs:=(uint32(time.Now().Unix()))
//	uid:=uint32(0)
//	randomInt:=uint32(58964981)
//	expiredTs:=uint32(unixTs+3600)
//
//	var publicSharingKey,sharingError = DynamicKey5.GeneratePublicSharingKey(appID, appCertificate, channelName, unixTs, randomInt, uid, expiredTs)
//	if sharingError == nil {
//		fmt.Println(publicSharingKey)
//	}
//
//	var mediaChannelKey,channelError = DynamicKey5.GenerateMediaChannelKey(appID, appCertificate, channelName, unixTs, randomInt, uid, expiredTs)
//	if channelError == nil {
//		fmt.Println(mediaChannelKey)
//	}
//
//	var recordingKey,recordingError = DynamicKey5.GenerateRecordingKey(appID, appCertificate, channelName, unixTs, randomInt, uid, expiredTs)
//	if recordingError == nil {
//		fmt.Println(recordingKey)
//	}
//
//	var noUploadKey,noUploadError = DynamicKey5.GenerateInChannelPermissionKey(appID, appCertificate, channelName, unixTs, randomInt, uid, expiredTs, DynamicKey5.NoUpload)
//	if noUploadError == nil {
//		fmt.Println(noUploadKey)
//	}
//
//	var audioVideoUploadKey,audioVideoUploadError = DynamicKey5.GenerateInChannelPermissionKey(appID, appCertificate, channelName, unixTs, randomInt, uid, expiredTs, DynamicKey5.AudioVideoUpload)
//	if audioVideoUploadError == nil {
//		fmt.Println(audioVideoUploadKey)
//	}
//}
var (
	count int32
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	channw()
}
func mut() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
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
