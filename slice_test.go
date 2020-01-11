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
	u   slice.UInter
	u8  slice.UInter8
	u16 slice.UInter16
	u32 slice.UInter32
	u64 slice.UInter64
	v   slice.Interfacer
)

func Test(t *testing.T) {

	var (
		s = slice.NewStringer()
	)
	s.Bounds(0)
}
