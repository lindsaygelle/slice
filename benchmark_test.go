package slice_test

import (
	"math/rand"
	"testing"

	"github.com/lindsaygelle/slice"
)

func BenchmarkAppend(b *testing.B) {
	slice := &slice.Slice[int]{}
	values := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		values[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		slice.Append(values[i])
	}
}

func BenchmarkAppendFunc(b *testing.B) {
	slice := &slice.Slice[int]{}
	values := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		values[i] = i
	}
	fn := func(i int, value int) bool {
		return true
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		slice.AppendFunc(values, fn)
	}
}

func BenchmarkAppendLength(b *testing.B) {
	slice := &slice.Slice[int]{}
	values := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		values[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.AppendLength(values...)
	}
}

func BenchmarkBounds(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.Bounds(i)
	}
}

func BenchmarkClone(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.Clone()
	}
}

func BenchmarkConcatenate(b *testing.B) {
	slice1 := &slice.Slice[int]{}
	slice2 := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice1.Append(i)
		slice2.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice1.Concatenate(slice2)
	}
}

func BenchmarkConcatenateFunc(b *testing.B) {
	slice1 := &slice.Slice[int]{}
	slice2 := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice1.Append(i)
		slice2.Append(i)
	}
	fn := func(i int, value int) bool {
		return true
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice1.ConcatenateFunc(slice2, fn)
	}
}

func BenchmarkConcatenateLength(b *testing.B) {
	slice1 := &slice.Slice[int]{}
	slice2 := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice1.Append(i)
		slice2.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice1.ConcatenateLength(slice2)
	}
}

func BenchmarkContains(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	value := rand.Intn(b.N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.Contains(value)
	}
}

func BenchmarkContainsMany(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	values := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		values[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.ContainsMany(values...)
	}
}

func BenchmarkDeduplicate(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(rand.Intn(b.N))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.Deduplicate()
	}
}

func BenchmarkDelete(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.Delete(i % slice.Length())
	}
}

func BenchmarkDeleteFunc(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	fn := func(i int, value int) bool {
		return true
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.DeleteManyFunc(fn)
	}
}

func BenchmarkDeleteLength(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.DeleteLength(i % slice.Length())
	}
}

func BenchmarkDeleteOK(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.DeleteOK(i % slice.Length())
	}
}

func BenchmarkDeleteUnsafe(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.DeleteUnsafe(i % slice.Length())
	}
}

func BenchmarkEach(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	fn := func(i int, value int) {}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.Each(fn)
	}
}

func BenchmarkEachBreak(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	fn := func(i int, value int) bool {
		return true
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.EachBreak(fn)
	}
}

func BenchmarkEachOK(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	fn := func(i int, value int) bool {
		return true
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.EachOK(fn)
	}
}

func BenchmarkEachReverse(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	fn := func(i int, value int) {}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.EachReverse(fn)
	}
}

func BenchmarkEachReverseBreak(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	fn := func(i int, value int) bool {
		return true
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.EachReverseBreak(fn)
	}
}

func BenchmarkEachReverseOK(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	fn := func(i int, value int) bool {
		return true
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.EachReverseOK(fn)
	}
}

func BenchmarkEqual(b *testing.B) {
	slice1 := &slice.Slice[int]{}
	slice2 := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice1.Append(i)
		slice2.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice1.Equal(slice2)
	}
}

func BenchmarkEqualFunc(b *testing.B) {
	slice1 := &slice.Slice[int]{}
	slice2 := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice1.Append(i)
		slice2.Append(i)
	}
	fn := func(i int, a, b int) bool {
		return a == b
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice1.EqualFunc(slice2, fn)
	}
}

func BenchmarkEqualLength(b *testing.B) {
	slice1 := &slice.Slice[int]{}
	slice2 := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice1.Append(i)
		slice2.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice1.EqualLength(slice2)
	}
}

func BenchmarkFetch(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.Fetch(i)
	}
}

func BenchmarkFetchLength(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = slice.FetchLength(i)
	}
}

func BenchmarkFilter(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	fn := func(i int, value int) bool {
		return value%2 == 0
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.Filter(fn)
	}
}

func BenchmarkFindIndex(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	fn := func(value int) bool {
		return value == b.N/2
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = slice.FindIndex(fn)
	}
}

func BenchmarkGet(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = slice.Get(i)
	}
}

func BenchmarkGetLength(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, _ = slice.GetLength(i)
	}
}

func BenchmarkIsEmpty(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.IsEmpty()
	}
}

func BenchmarkIsPopulated(b *testing.B) {
	slice := &slice.Slice[int]{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.IsPopulated()
	}
}

func BenchmarkLength(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.Length()
	}
}

func BenchmarkMap(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	fn := func(i int, value int) int {
		return value * value
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.Map(fn)
	}
}

func BenchmarkPop(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.Pop()
	}
}

func BenchmarkPopLength(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = slice.PopLength()
	}
}

func BenchmarkReduce(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	fn := func(i, acc, value int) int {
		return acc + value
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.Reduce(fn)
	}
}

func BenchmarkReverse(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.Reverse()
	}
}

func BenchmarkShuffle(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.Shuffle()
	}
}

func BenchmarkSlice(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	start := b.N / 4
	end := 3 * b.N / 4
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.Slice(start, end)
	}
}

func BenchmarkSortFunc(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(rand.Intn(b.N))
	}
	fn := func(i, j, a, b int) bool {
		return a < b
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slice.SortFunc(fn)
	}
}

func BenchmarkSwap(b *testing.B) {
	slice := &slice.Slice[int]{}
	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
	index1 := b.N / 4
	index2 := 3 * b.N / 4
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		slice.Swap(index1, index2)
	}
}
