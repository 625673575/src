package tf

import (
	"testing"
	"reflect"
	"unsafe"
	"math"
)

func Benchmark_BubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
func Test_NewTensor(t *testing.T) {
	vs := [] float64{1.23, 12.3, 111, 0.98, 0.34, 0.3333}
	tensor := New(vs, []int{2, 3}, false)
	if reflect.DeepEqual(vs, tensor.data) {
		t.Log("pass")
	} else {
		t.Error("fail")
	}
	t.Log(tensor.data, tensor.shape, tensor.step)
}
func Test_Linear(t *testing.T) {
	tensor := Linear(2.34, 8.0, 2.2)
	t.Log(tensor, tensor.shape)
}

func Test_Ceil(t *testing.T) {
	tensor := Linear(2.34, 1333.0, 2.2)
	println(tensor)
	newtensor := Ceil(tensor)
	println((len(tensor.data)), unsafe.Offsetof(tensor.shape), unsafe.Offsetof(tensor.step))
	t.Log(newtensor)
	tensor.Ceil()
	println(tensor)
	println(newtensor)
	t.Log(tensor)
}

func Test_Add(t *testing.T) {
	tensor := Linear(math.MinInt32, math.MaxInt32, 1024)
	tensor2 := Linear(math.MinInt32, math.MaxInt32, 1024)
	tensor.Add(tensor2)
	t.Log(tensor.Size())
}
func Benchmark_Add(b *testing.B) {
	tensor := Linear(math.MinInt32, math.MaxInt32, 1024)
	tensor2 := Linear(math.MinInt32, math.MaxInt32, 1024)
	for i := 0; i < b.N; i++ {
		tensor.Add(tensor2)
	}
}
