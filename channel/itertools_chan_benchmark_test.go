package channel

import (
	"fmt"
	"math"
	"testing"
)

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := Count(int8(0), int8(1))
		for a := range Iter(m) {
			a = a
		}
	}

}
func BenchmarkLoop(b *testing.B) {
	var a interface{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < math.MaxInt8; j++ {
			a = j
		}
	}
	fmt.Println(a)
}
