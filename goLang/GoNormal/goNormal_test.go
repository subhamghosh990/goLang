package normal

import (
	"testing"
)

func BenchmarkPrintData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := GetData(i)
		PrintData(res)
	}
}

func FuzzData(f *testing.F) {
	//f.Add(100)
	f.Fuzz(func(t *testing.T, in int) {
		res := GetData(in)
		PrintData(res)
	})
}

func FuzzCompareDat(f *testing.F) {
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		CompareData(a, b)
	})
}
