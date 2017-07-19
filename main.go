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
func mut(){
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
}
func chann(){
	ch := make(chan int)
	go func() {
		var sum int = 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		ch <- sum
	}()

	fmt.Println(<-ch)
}
func channw(){
	one := make(chan int)
	two := make(chan int)
	go func() {
		v:=<-one
		two<-v

	}()
	time.Sleep(time.Second)
	go func() {
		one<-100
	}()
	fmt.Println(<-two)

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
