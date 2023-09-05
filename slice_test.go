package slice_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/gellel/slice"
)

var (
	b    slice.Byte       // []byte
	c64  slice.Complex64  // []complex64
	c128 slice.Complex128 // []complex128
	f32  slice.Float32    // []float32
	f64  slice.Float64    // []float64
	i    slice.Int        // []interface{}
	i8   slice.Int8       // []int8
	i16  slice.Int16      // []int16
	i32  slice.Int32      // []int32
	i64  slice.Int64      // []int64
	r    slice.Rune       // []rune
	s    *slice.Slice     // []interface{}
	u    slice.UInt       // []uint
	u8   slice.UInt8      // []uint8
	u16  slice.UInt16     // []uint16
	u32  slice.UInt32     // []uint32
	u64  slice.UInt64     // []uint64
	v    slice.Interface  // []interface{}
)

func Test(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	s = &slice.Slice{}
}

func TestAppend(t *testing.T) {
	if ok := (*s.Append("a"))[0].(string) == "a"; !ok {
		t.Fatalf("(&slice.Slice.Append(interface{})) != (interface{}))")
	}
}

func TestBounds(t *testing.T) {
	if ok := s.Bounds(s.Len() - 1); !ok {
		t.Fatalf("(&slice.Slice.Bounds(int) bool) != true")
	}
}

func TestConcatenate(t *testing.T) {
	if ok := (*s.Concatenate(&slice.Slice{"b"}))[1].(string) == "b"; !ok {
		t.Fatalf("(&slice.Slice.Concatenate(interface{})) != (interface{}))")
	}
	var (
		n = len(*s)
	)
	if ok := (len(*s.Concatenate(nil))) == n; !ok {
		t.Fatalf("(&slice.Concatenate(nil).Len()) != nil")
	}
}

func TestDelete(t *testing.T) {
	s.Append("c")
	var (
		n = s.Len()
		z = float64(s.Len()) / 2
	)
	var (
		mid = (*s)[int(z)]
	)
	if ok := s.Delete(int(z)).Len() != n; !ok {
		t.Fatalf("(&slice.Slice.Delete(int)) != true")
	}
	for _, v := range *s {
		if ok := (v.(string)) != mid; !ok {
			t.Fatalf("(&slice.Slice.Delete(int)) != true")
		}
	}
}

func TestEach(t *testing.T) {
	var (
		n int
	)
	s.Each(func(i int, v interface{}) {
		if ok := (*s)[i] == v; !ok {
			t.Fatalf("(&slice.Slice.Each(int, interface{})) != (interface{})")
		}
		if ok := i == n; !ok {
			t.Fatalf("(&slice.Slice.Each(i int, interface{})) != i")
		}
		n = n + 1
	})
}

func TestEachBreak(t *testing.T) {
	var (
		n int
	)
	s.EachBreak(func(i int, _ interface{}) bool {
		n = i
		return false
	})
	if ok := n == 0; !ok {
		t.Fatalf("(&slice.Slice.EachBreak(int, interface{}) bool) != true")
	}
}

func TestEachReverse(t *testing.T) {
	var (
		n = s.Len() - 1
	)
	s.EachReverse(func(i int, v interface{}) {
		if ok := (*s)[i] == v; !ok {
			t.Fatalf("(&slice.Slice.EachReverse(int, interface{})) != (interface{})")
		}
		if ok := i == n; !ok {
			t.Fatalf("(&slice.Slice.EachReverse(i int, interface{})) != i")
		}
		n = n - 1
	})
}

func TestEachReverseBreak(t *testing.T) {
	var (
		n int
	)
	s.EachReverseBreak(func(i int, _ interface{}) bool {
		n = i
		return false
	})
	if ok := n == s.Len()-1; !ok {
		t.Fatalf("(&slice.Slice.EachReverseBreak(int, interface{}) bool) != true")
	}
}

func TestFetch(t *testing.T) {
	if ok := s.Fetch(s.Len()+1) == nil; !ok {
		t.Fatalf("(&slice.Slice.Fetch(int) interface{}) != true")
	}
	if ok := s.Fetch(0) != nil; !ok {
		t.Fatalf("(&slice.Slice.Fetch(int) interface{}) != true")
	}
}

func TestGet(t *testing.T) {
	if _, ok := s.Get(0); !ok {
		t.Fatalf("(&slice.Slice.Get(int) (_, bool)) != true")
	}
	if v, _ := s.Get(0); v != (*s)[0] {
		t.Fatalf("(&slice.Slice.Get(int) (interface{}, _) != interface{}")
	}
}

func TestMake(t *testing.T) {
	if ok := s.Make(10).Len() == 10; !ok {
		t.Fatalf("(&slice.Make(int).Len()) != n")
	}
}

func TestMakeEach(t *testing.T) {
	var (
		v = []interface{}{1, 2, 3, 4, 5}
	)
	if ok := s.MakeEach(v...).Len() == len(v); !ok {
		t.Fatalf("(&slice.MakeEach(...interface{}).Len()) != n")
	}
	s.Each(func(i int, x interface{}) {
		if ok := v[i] == x; !ok {
			t.Fatalf("(&slice.MakeEach(...interface{})) != interface{}")
		}
	})
}

func TestMakeEachReverse(t *testing.T) {
	var (
		v = []interface{}{1, 2, 3, 4, 5}
	)
	if ok := s.MakeEachReverse(v...).Len() == len(v); !ok {
		t.Fatalf("(&slice.MakeEachReverse(...interface{}).Len()) != n")
	}
	s.EachReverse(func(i int, x interface{}) {
		if ok := v[i] == x; !ok {
			t.Fatalf("(&slice.MakeEachReverse(...interface{})) != interface{}")
		}
	})
}

func TestMap(t *testing.T) {
	var (
		x = []int{}
	)
	s.Each(func(_ int, v interface{}) {
		x = append(x, v.(int)*2)
	})
	s.Map(func(i int, v interface{}) interface{} {
		var x = v.(int)
		x = x * 2
		return x
	})
	s.Each(func(i int, v interface{}) {
		if ok := x[i] == v.(int); !ok {
			t.Fatalf("(&slice.Map(func(int, interface{}) interface{}})) != interface{}")
		}
	})
}

func TestPoll(t *testing.T) {
	var (
		n = rand.Intn(100)
	)
	if n == 0 {
		n = 1
	}
	var (
		v = make([]interface{}, n)
	)
	for i := range v {
		v[i] = rand.Intn(100)
	}
	s.MakeEach(v...)
	var (
		x = s.Poll()
	)
	if ok := x == v[0]; !ok {
		t.Fatalf("(&slice.Poll() interface{}) != interface{}")
	}
	if ok := len(v) != s.Len(); !ok {
		t.Fatalf("(&slice.Poll() interface{}); (&slice.Len()) == len(v)")
	}
	for i := s.Len(); i > 0; i-- {
		x = s.Poll()
		if ok := x != nil; !ok {
			t.Fatalf("(&slice.Poll() interface{}) != interface{}")
		}
		if ok := x == v[len(v)-i]; !ok {
			t.Fatalf("(&slice.Poll() interface{}) != []interface{}[i]")
		}
	}
	if ok := s.Len() == 0; !ok {
		t.Fatalf("(&slice.Poll() interface{}); (&slice.Len()) != 0")
	}
}

func TestPop(t *testing.T) {
	var (
		n = rand.Intn(100)
	)
	if n == 0 {
		n = 1
	}
	var (
		v = make([]interface{}, rand.Intn(100))
	)
	for i := range v {
		v[i] = rand.Intn(100)
	}
	s.MakeEach(v...)
	var (
		x = s.Pop()
	)
	if ok := x == v[len(v)-1]; !ok {
		t.Fatalf("(&slice.Pop() interface{}) != interface{}")
	}
	if ok := len(v) != s.Len(); !ok {
		t.Fatalf("(&slice.Pop() interface{}); (&slice.Len()) == len(v)")
	}
	for i := s.Len(); i > 0; i-- {
		x = s.Pop()
		if ok := x != nil; !ok {
			t.Fatalf("(&slice.Pop() interface{}) != interface{}")
		}
		if ok := x == v[i-1]; !ok {
			t.Fatalf("(&slice.Pop() interface{}) != []interface{}[i]")
		}
	}
	if ok := s.Len() == 0; !ok {
		t.Fatalf("(&slice.Pop() interface{}); (&slice.Len()) != 0")
	}
}

func TestPrecatenate(t *testing.T) {
	var (
		head = 1 + rand.Intn(10-1)
		tail = head + rand.Intn(20-head)
	)
	s = &slice.Slice{}
	s.Append(head)
	s.Precatenate((&slice.Slice{}).Append(tail))
	if ok := s.Len() == 2; !ok {
		t.Fatalf("(&slice.Precatenate(&slice.Slice{}).Len()) != n")
	}
	if ok := s.Fetch(0) == tail; !ok {
		t.Fatalf("(&slice.Precatenate(&slice.Slice{}).Fetch(0) != tail")
	}
	if ok := s.Fetch(1) == head; !ok {
		t.Fatalf("(&slice.Precatenate(&slice.Slice{}).Fetch(1) != head")
	}
	s.Precatenate(nil)
	if ok := s.Len() == 2; !ok {
		t.Fatalf("(&slice.Precatenate(nil).Len()) != nil")
	}
}

func TestPrepend(t *testing.T) {
	var (
		head = 1 + rand.Intn(10-1)
		tail = head + rand.Intn(20-head)
	)
	s = &slice.Slice{tail}
	s.Prepend(head)
	if ok := s.Len() == 2; !ok {
		t.Fatalf("(&slice.Prepend(interface{}).Len()) != n")
	}
	if ok := s.Fetch(0) == head; !ok {
		t.Fatalf("(&slice.Prepend(interface{}).Fetch(0) != head")
	}
	if ok := s.Fetch(1) == tail; !ok {
		t.Fatalf("(&slice.Prepend(interface{}).Fetch(1) != tail")
	}
	s.Prepend()
	if ok := s.Len() == 2; !ok {
		t.Fatalf("(&slice.Prepend(nil).Len()) != nil")
	}
}

func TestPush(t *testing.T) {
	var (
		n = rand.Intn(100)
	)
	if n == 0 {
		n = 1
	}
	var (
		v = make([]interface{}, rand.Intn(100))
	)
	for i := range v {
		v[i] = rand.Intn(100)
	}
	s = &slice.Slice{}
	s.Push(v...)
	if ok := s.Len() == len(v); !ok {
		t.Fatalf("(&slice.Push(interface{}...).Len()) != n")
	}
	s.Each(func(i int, x interface{}) {
		if ok := v[i] == x; !ok {
			t.Fatalf("(&slice.Push([]interface{}...).Fetch(i) interface{}) != []interface{}[i]")
		}
	})
}

func TestReplace(t *testing.T) {
	var (
		a = "a"
		b = "b"
	)
	s = &slice.Slice{a}
	s.Replace(0, b)
	if ok := (*s)[0].(string) == b; !ok {
		t.Fatalf("(&slice.Replace(int, interface{}).Fetch(i)) != interface{}")
	}
}
