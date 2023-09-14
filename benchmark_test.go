package slice_test

import (
	"testing"

	"github.com/lindsaygelle/slice"
)

func BenchmarkAppend(b *testing.B) {
	s := &slice.Slice[int]{1, 2, 3, 4, 5}

	// Run the benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Append(6)
	}
}

func BenchmarkAppendLength(b *testing.B) {
	s := &slice.Slice[int]{1, 2, 3, 4, 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.AppendLength(6)
	}
}

func BenchmarkConcatenate(b *testing.B) {
	s1 := &slice.Slice[int]{}
	s2 := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		s1.Append(i)
		s2.Append(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s1.Concatenate(s2)
	}
}

func BenchmarkDelete(b *testing.B) {
	s := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		s.Append(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Delete(i % s.Length())
	}
}

func BenchmarkEach(b *testing.B) {
	s := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		s.Append(i)
	}
	fn := func(_ int, _ int) {}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Each(fn)
	}
}

func BenchmarkFindIndex(b *testing.B) {
	s := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		s.Append(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = s.FindIndex(func(value int) bool {
			return value == i
		})
	}
}

func BenchmarkMap(b *testing.B) {
	s := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		s.Append(i)
	}
	fn := func(_ int, value int) int {
		return value * 2
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Map(fn)
	}
}

func BenchmarkReverse(b *testing.B) {
	s := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		s.Append(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Reverse()
	}
}

func BenchmarkSet(b *testing.B) {
	s := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		s.Append(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Set()
	}
}
