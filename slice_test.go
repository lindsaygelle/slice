package slice_test

import (
	"testing"

	"github.com/gellel/slice"
)

var (
	b   slice.Byter
	f32 slice.Floater32
	f64 slice.Floater64
	i   slice.Inter
	i8  slice.Inter8
	i16 slice.Inter16
	i32 slice.Inter32
	i64 slice.Inter64
	s   *slice.Slice
	u   slice.UInter
	u8  slice.UInter8
	u16 slice.UInter16
	u32 slice.UInter32
	u64 slice.UInter64
	v   slice.Interfacer
)

func Test(t *testing.T) {
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
