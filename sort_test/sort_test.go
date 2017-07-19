package sort_test

import (
	"testing"
	"reflect"
	"math/rand"
	"runtime"
)

var x =12
var y =1
var tests=[]int{12,32343,2,234,1223,35, 223,235,566,66,542,9}
var rights=[]int{2,9 ,12, 35 ,66 ,223 ,234, 235, 542 ,566 ,1223, 32343}

func swap(x,y *int){
	*x,*y=*y,*x
}
func BubbleSort( list []int)[]int {
	l:=len(list)
	for i:=0;i<l;i++{
		for j:=i;j<l;j++{
			if list[i]>list[j]{
			swap(	&list[i],&list[j])
			}
		}
	}
	return list
}
func QuickSort(list []int)(ret []int){
	l:=len(list)
	var low,mid,high []int
	if l<2{
		return list
	}else{
		ran:=list[0]
		for _,v:= range list{
			if v<ran{
				low=append(low,v)
			}else if v== ran{
				mid=append(mid,v)
			}else{
				high=append(high,v)
			}
		}
		ret=append(ret, QuickSort(low)...)
		ret=append(ret,mid...)
		ret=append(ret,QuickSort(high)...)
		return ret
	}
}
func Test_BubbleSort(t *testing.T){
	x:=BubbleSort(tests)
	if reflect.DeepEqual(x,rights){
		t.Log("pass",x)
	}else{
		t.Error("fail",x)
	}
}
func Benchmark_BubbleSort(b *testing.B) {
	for i :=0;i<b.N;i++{
		BubbleSort(tests)
	}
}
func Test_QuickSort(t *testing.T){
	x:=qsort(tests)
	if reflect.DeepEqual(x,rights){
		t.Log("pass",x)
	}else{
		t.Error("fail",x)
	}
}
var worker=runtime.NumCPU()
func qsort(a []int) []int {
	if len(a) < 2 { return a }

	left, right := 0, len(a) - 1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	qsort(a[:left])
	qsort(a[left + 1:])


	return a
}
func Benchmark_QuickSort(b *testing.B) {
	for i :=0;i<b.N;i++{
		qsort(tests)
	}
}