package channel

import (
	"math"
	"reflect"

	"runtime"

	"golang.org/x/tools/container/intsets"
)

type IterMgr struct {
	c      chan interface{}
	signal chan interface{}
}

const (
	Done = 0
)

func newIterMgr() *IterMgr {
	iter := &IterMgr{
		c:      make(chan interface{}),
		signal: make(chan interface{}),
	}
	runtime.SetFinalizer(iter, Finalizer)
	return iter
}

func New(items ...interface{}) *IterMgr {
	iter := newIterMgr()
	go func() {
		flag := false
		for _, i := range items {
			select {
			case iter.c <- i:
			case <-iter.signal:
				flag = true
				goto close
			}
		}
	close:
		iter.closeChan(flag)
	}()
	return iter
}

func Finalizer(i *IterMgr) {
	i.signal <- Done
}

func (i *IterMgr) Next() interface{} {
	res, isOpen := <-i.c
	if !isOpen {
		panic("StopIteration")
	}
	return res
}

func (i *IterMgr) closeChan(f bool) {
	close(i.c)
	if !f {
		<-i.signal
	}
	close(i.signal)
}

func Iter(i *IterMgr) chan interface{} {
	return i.c
}

func Count(start interface{}, step interface{}) *IterMgr {
	if reflect.TypeOf(start) != reflect.TypeOf(step) {
		panic("params must be the same type")
	}
	switch start.(type) {
	case int:
		return countInt(start.(int), step.(int))
	case int8:
		return countInt8(start.(int8), step.(int8))
	case int16:
		return countInt16(start.(int16), step.(int16))
	case int32:
		return countInt32(start.(int32), step.(int32))
	case int64:
		return countInt64(start.(int64), step.(int64))
	case float32:
		return countFloat32(start.(float32), step.(float32))
	case float64:
		return countFloat64(start.(float64), step.(float64))
	}
	return nil
}

func countInt(start int, step int) *IterMgr {
	iter := newIterMgr()
	go func() {
		flag := false
		for i := start; i < intsets.MaxInt; i += step {
			select {
			case iter.c <- i:
			case <-iter.signal:
				flag = true
				goto close
			}
		}
	close:
		iter.closeChan(flag)
	}()
	return iter
}

func countInt64(start int64, step int64) *IterMgr {
	iter := newIterMgr()
	go func() {
		flag := false
		for i := start; i < math.MaxInt64; i += step {
			select {
			case iter.c <- i:
			case <-iter.signal:
				flag = true
				goto close

			}
		}
	close:
		iter.closeChan(flag)
	}()
	return iter
}

func countInt16(start int16, step int16) *IterMgr {
	iter := newIterMgr()
	go func() {
		flag := false
		for i := start; i < math.MaxInt16; i += step {
			select {
			case iter.c <- i:
			case <-iter.signal:
				flag = true
				goto close
			}
		}
	close:
		iter.closeChan(flag)
	}()
	return iter
}

func countInt32(start int32, step int32) *IterMgr {
	iter := newIterMgr()
	go func() {
		flag := false
		for i := start; i < math.MaxInt32 && !flag; i += step {
			select {
			case iter.c <- i:
			case <-iter.signal:
				flag = true
				goto close
			}
		}
	close:
		iter.closeChan(flag)
	}()
	return iter
}

func countInt8(start int8, step int8) *IterMgr {
	iter := newIterMgr()
	go func() {
		flag := false
		for i := start; i < math.MaxInt8; i += step {
			select {
			case iter.c <- i:
			case <-iter.signal:
				flag = true
				goto close
			}
		}
	close:
		iter.closeChan(flag)
	}()
	return iter
}

func countFloat32(start float32, step float32) *IterMgr {
	iter := newIterMgr()
	go func() {
		flag := false
		for i := start; i < math.MaxFloat32; i += step {
			select {
			case iter.c <- i:
			case <-iter.signal:
				flag = true
				goto close
			}
		}
	close:
		iter.closeChan(flag)
	}()
	return iter
}

func countFloat64(start float64, step float64) *IterMgr {
	iter := newIterMgr()
	go func() {
		flag := false
		for i := start; i < math.MaxFloat64; i += step {
			select {
			case iter.c <- i:
			case <-iter.signal:
				flag = true
				goto close
			}
		}
	close:
		iter.closeChan(flag)
	}()
	return iter
}

func Cycle(item interface{}) *IterMgr {
	switch item.(type) {
	case string:
		return cycString(item.(string))
	case *IterMgr:
		return cycIterMgr(item.(*IterMgr))
	}
	if slice, ok := takeSliceArg(item); ok {
		return cycSlice(slice)
	}
	if smap, ok := takeMapArg(item); ok {
		return cycMap(smap)
	}
	return nil

}

func cycString(item string) *IterMgr {
	iter := newIterMgr()
	go func() {
		flag := false
		for {
			for _, i := range item {
				select {
				case iter.c <- string(i):
				case <-iter.signal:
					flag = true
					goto close
				}
			}
		}
	close:
		iter.closeChan(flag)
	}()
	return iter
}

func cycIterMgr(item *IterMgr) *IterMgr {
	iter := newIterMgr()
	temp := make([]interface{}, 0)
	for i := range Iter(item) {
		temp = append(temp, i)
	}
	go func() {
		flag := false
		for {
			for _, i := range temp {
				select {
				case iter.c <- i:
				case <-iter.signal:
					flag = true
					goto close
				}
			}
		}
	close:
		iter.closeChan(flag)
	}()
	return iter
}

func cycSlice(item []interface{}) *IterMgr {
	iter := newIterMgr()
	go func() {
		flag := false
		for {
			for _, i := range item {
				select {
				case iter.c <- i:
				case <-iter.signal:
					flag = true
					goto close
				}
			}
		}
	close:
		iter.closeChan(flag)
	}()
	return iter
}

func cycMap(item map[string]interface{}) *IterMgr {
	iter := newIterMgr()
	go func() {
		flag := false
		for {
			for i := range item {
				select {
				case iter.c <- i:
				case <-iter.signal:
					flag = true
					goto close
				}
			}
		}
	close:
		iter.closeChan(flag)
	}()
	return iter
}

func DropWhile(lambda func(interface{}) bool, i *IterMgr) *IterMgr {
	iter := newIterMgr()
	go func() {
		flag := false
		for i := range Iter(i) {
			if lambda(i) {
				select {
				case iter.c <- i:
				case <-iter.signal:
					flag = true
					goto close
				}
			}
		}
	close:
		iter.closeChan(flag)
	}()
	return iter
}

func Repeat(item interface{}, times int) *IterMgr {
	iter := newIterMgr()
	go func() {
		flag := false
		for i := 0; i < times; i++ {
			select {
			case iter.c <- item:
			case <-iter.signal:
				flag = true
				goto close
			}
		}
	close:
		iter.closeChan(flag)
	}()
	return iter
}

func Imap(f func(a interface{}, b interface{}) interface{}, aIter *IterMgr, bIter *IterMgr) *IterMgr {
	iter := newIterMgr()
	go func() {
		flag := false
		for value := range Iter(aIter) {
			select {
			case iter.c <- f(value, bIter.Next()):
			case <-iter.signal:
				flag = true
				goto close
			}
		}
	close:
		iter.closeChan(flag)
	}()
	return iter
}
