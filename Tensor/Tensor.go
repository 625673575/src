package tf

import (
	"unsafe"
	"reflect"
	"fmt"
	"math"
	"runtime"
	"sync"
)

type DTYPE int

const (
	DTYPE_FLOAT64 = 0
)

type Tensor struct {
	data  []float64
	shape []int
	step  int
}

var CORE_NUM = runtime.NumCPU()

func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}
func (t *Tensor) Size() int {
	count := int(1)
	for _, v := range t.shape {
		count *= v
	}
	return count
}
func (t *Tensor) Rank() int {
	return len(t.shape)
}

func output(array []string) string {
	len := len(array)
	s := "["
	for i := 0; i < len-1; i++ {
		s += fmt.Sprint(array[i], ",")
	}
	s += fmt.Sprint(array[len-1], "]")
	return s
}

var currentRank = -1

func recursiveOutput(array []string, shape []int) []string {
	size := len(array)
	if currentRank < 0 {
		currentRank = len(shape) - 1
	}
	if size == 1 {
		currentRank = -1
		return array
	} else {
		start := 0
		width := shape[currentRank]
		s0 := make([]string, 0, 16)
		for i := 0; i < size; i++ {
			if (i-start)%width == 0 {
				sff := make([]string, width)
				for k := 0; k < width; k++ {
					sff[k] = fmt.Sprint(array[start+k])
				}
				start += width
				s0 = append(s0, output(sff))
			}
		}
		currentRank--
		return recursiveOutput(s0, shape)
	}
}
func (t *Tensor) StringArray() []string {
	s := make([]string, len(t.data))
	for i, v := range t.data {
		s[i] = fmt.Sprint(v)
	}
	return s
}
func (t *Tensor) String() string {
	s := recursiveOutput(t.StringArray(), t.shape)
	return s[0]
}
func New(data []float64, shape []int, ignoreSizeCheck bool) *Tensor {
	v := reflect.ValueOf(data)
	len := v.Len()
	if !ignoreSizeCheck {
		count := int(1)
		for _, v := range shape {
			count *= v
		}
		if len != count {
			panic("the shape is not match the data size")
		}
	}
	ret := new(Tensor)
	ret.shape = shape
	ret.data = data
	ret.step = int(unsafe.Sizeof(v)) / len
	return ret
}
func Zero(shape []int) *Tensor {
	count := int(1)
	for _, v := range shape {
		count *= v
	}
	data := make([]float64, count)
	ret := new(Tensor)
	ret.shape = shape
	ret.data = data
	ret.step = 8
	return ret
}

func Linear(start float64, end float64, gap float64) *Tensor {
	count := (int)(math.Ceil((end - start) / gap))
	r := make([]float64, count)
	c := 0
	for i := start; i <= end; i += gap {
		r[c] = i
		c++
	}
	t := new(Tensor)
	t.data = r
	t.shape = []int{count}
	t.step = 8
	return t
}
func ZeroBy(t *Tensor) *Tensor {
	tensor := new(Tensor)
	tensor.data = make([]float64, t.Size())
	tensor.shape = t.shape
	tensor.step = t.step
	return tensor
}
func Ceil(t *Tensor) *Tensor {
	tensor := ZeroBy(t)
	for i := 0; i < t.Size(); i++ {
		tensor.data[i] = math.Ceil(t.data[i])
	}
	return tensor
}
func (t *Tensor) Ceil() {
	for i := 0; i < t.Size(); i++ {
		t.data[i] = math.Ceil(t.data[i])
	}
}

var wg sync.WaitGroup

func (t *Tensor) GoCompute(arg *Tensor,fn func (arg0 *float64, arg1 *float64)) {
	split := (t.Size())/CORE_NUM + 1
	wg.Add(CORE_NUM)
	for k := 0; k < CORE_NUM; k++ {
		go func(n int) {
			start := n * split
			end := (n + 1) * split
			if end > t.Size() {
				end = t.Size()
			}
			for i := start; i < end; i++ {
				fn(&t.data[i],&arg.data[i])
			}
			wg.Done()
		}(k)
	}
	wg.Wait()
}
func (t *Tensor) NAdd(arg *Tensor){
	for i:=0;i<t.Size();i++{
		t.data[i]+=arg.data[i]
	}
}
func (t *Tensor) Add(arg *Tensor) {
	t.GoCompute(arg,addAtom)
}
func (t *Tensor) Sub(arg *Tensor) {
	t.GoCompute(arg,subAtom)
}
func addAtom(arg0 *float64, arg1 *float64) {
	*arg0 = *arg0 + *arg1
}
func subAtom(arg0 *float64, arg1 *float64) {
	*arg0 = *arg0 - *arg1
}
