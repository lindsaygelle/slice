// TestAppend tests Slice.Append.
func TestAppend(t *testing.T) {
	s := &slice.Slice[int]{}
	s.Append(1)
	if ok := len(*s) == 1; !ok {
		t.Fatalf("len(*Slice) != 1")
	}
	s = &slice.Slice[int]{}
	s.Append(1, 2)
	if ok := len(*s) == 2; !ok {
		t.Fatalf("len(*Slice) != 2")
	}
}

// TestAppendLength tests Slice.AppendLength.
func TestAppendLength(t *testing.T) {
	s := &slice.Slice[int]{}
	if n := s.AppendLength(1, 2, 3, 4); n != len(*s) {
		t.Fatalf("len(*Slice) != 4")
	}
}

// TestBounds tests Slice.Bounds.
func TestBounds(t *testing.T) {
	s := &slice.Slice[int]{}
	if ok := s.Bounds(0); ok {
		t.Fatalf("*Slice.Bounds() != false")
	}
	s.Append(1)
	if ok := s.Bounds(0); !ok {
		t.Fatalf("*Slice.Bounds() != true")
	}
}

// TestConcatenate tests Slice.Concatenate.
func TestConcatenate(t *testing.T) {
	s := &slice.Slice[int]{}
	s.Append(1)
	s.Concatenate(&slice.Slice[int]{2, 3})
	if ok := (*s)[0] == 1; !ok {
		t.Fatalf("*Slice[0] != 1")
	}
	if ok := (*s)[1] == 2; !ok {
		t.Fatalf("*Slice[1] != 2")
	}
	if ok := (*s)[2] == 3; !ok {
		t.Fatalf("*Slice[2] != 3")
	}
}

// TestConcatenateLength tests Slice.ConcatenateLength.
func TestConcatenateLength(t *testing.T) {
	s := &slice.Slice[int]{}
	if n := s.ConcatenateLength(&slice.Slice[int]{1}); n != len(*s) {
		t.Fatalf("len(*Slice) != %d", len(*s))
	}
}

// TestDelete tests Slice.Delete.
func TestDelete(t *testing.T) {
	s := &slice.Slice[int]{1}
	s.Delete(0)
	if ok := len(*s) == 0; !ok {
		t.Fatalf("len(*Slice) != 0")
	}
}

// TestDeleteLength tests Slice.DeleteLength.
func TestDeleteLength(t *testing.T) {
	s := &slice.Slice[int]{1}
	if n := s.DeleteLength(0); n != len(*s) {
		t.Fatalf("len(*Slice) != 0")
	}
}

// TestEach tests Slice.Each.
func TestEach(t *testing.T) {
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	s.Each(func(i int, value int) {
		if ok := (*s)[i] == value; !ok {
			t.Fatalf("*Slice[%d] != %d", i, value)
		}
	})
}

// TestEachBreak tests Slice.EachBreak.
func TestEachBreak(t *testing.T) {
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	count := 0
	s.EachBreak(func(i int, value int) bool {
		count = count + 1
		return false
	})
	if ok := count == 1; !ok {
		t.Fatalf("count != 1")
	}
}

// TestEachReverse tests Slice.EachReverse.
func TestEachReverse(t *testing.T) {
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	s.EachReverse(func(i int, value int) {
		if ok := (*s)[i] == value; !ok {
			t.Fatalf("*Slice[%d] != %d", i, value)
		}
	})
}

// TestEachReverseBreak tests Slice.EachReverseBreak.
func TestEachReverseBreak(t *testing.T) {
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	count := 0
	s.EachReverseBreak(func(i int, value int) bool {
		count = count + 1
		return false
	})
	if ok := count == 1; !ok {
		t.Fatalf("count != 1")
	}
}

// TestFetch tests Slice.Fetch.
func TestFetch(t *testing.T) {
	s := &slice.Slice[int]{1}
	for i, _ := range (*s) {
		value := s.Fetch(i)
		if ok := value == (*s)[i]; !ok {
			t.Fatalf("%d != %d", value, (*s)[i])
		}
	}
	// Deliberately empty and check default is returned.
	s = &slice.Slice[int]{}
	value = s.Fetch(1)
	if ok := value == 0; !ok {
		t.Fatalf("%d != 0", value)
	}
}

// TestFetchLength tests Slice.FetchLength.
func TestFetchLength(t *testing.T) {
	s := &slice.Slice[int]{1, 2}
	for i, _ := range (*s) {
		_, value := s.FetchLength(i)
		if ok := value == len(*s); !ok {
			t.Fatalf("%d != %d", value, len(*s))
		}
	}
}

// TestGet tests Slice.Get.
func TestGet(t *testing.T) {
	s := &slice.Slice[int]{1}
	for i, _ := range (*s) {
		value, ok := s.Get(i)
		if value != (*s)[i] {
			t.Fatalf("%d != %d", value, (*s)[i])
		}
		if !ok {
			t.Fatalf("%t != true", ok)
		}
	}
}

// TestGetLength tests Slice.GetLength.
func TestGetLength(t *testing.T) {
	s := &slice.Slice[int]{1}
	for i, _ := range (*s) {
		value, length, ok := s.GetLength(i)
		if value != (*s)[i] {
			t.Fatalf("%d != %d", value, (*s)[i])
		}
		if length != len(*s) {
			t.Fatalf("%d = %d", length, len(*s))
		}
		if !ok {
			t.Fatalf("%t != true", ok)
		}
	}
}

// TestLength tests Slice.Length.
func TestLength(t *testing.T) {
	s := &slice.Slice[int]{}	
	if ok := s.Length() == len(*s); !ok {
		t.Fatalf("len(*Slice) != %d", len(*s))
	}
}

// TestMake tests Slice.Make.
func TestMake(t *testing.T) {
	s := &slice.Slice[int]{}
	s.Make()
}

// TestMakeEach tests Slice.MakeEach.
func TestMakeEach(t *testing.T) {
	s := &slice.Slice[int]{}
	s.MakeEach()
}

// TestMakeEachReverse tests Slice.MakeEachReverse.
func TestMakeEachReverse(t *testing.T) {
	s := &slice.Slice[int]{}
	s.MakeEachReverse()
}

// TestMap tests Slice.Map.
func TestMap(t *testing.T) {
	s := &slice.Slice[int]{}
	s.Map()
}

// TestPoll tests Slice.Poll.
func TestPoll(t *testing.T) {
	s := &slice.Slice[int]{}
	s.Poll()
}

// TestPollLength tests Slice.PollLength.
func TestPollLength(t *testing.T) {
	s := &slice.Slice[int]{}
	s.PollLength()
}

// TestPollOK tests Slice.PollOK.
func TestPollOK(t *testing.T) {
	s := &slice.Slice[int]{}
	s.PollOK()
}

// TestPop tests Slice.Pop.
func TestPop(t *testing.T) {
	s := &slice.Slice[int]{}
	s.Pop()
}

// TestPopLength tests Slice.PopLength.
func TestPopLength(t *testing.T) {
	s := &slice.Slice[int]{}
	s.PopLength()
}

// TestPopOK tests Slice.PopOK.
func TestPopOK(t *testing.T) {
	s := &slice.Slice[int]{}
	s.PopOK()
}

// TestPrecatenate tests Slice.Precatenate.
func TestPrecatenate(t *testing.T) {
	s := &slice.Slice[int]{}
	s.Precatenate()
}

// TestPrecatenateLength tests Slice.PrecatenateLength.
func TestPrecatenateLength(t *testing.T) {
	s := &slice.Slice[int]{}
	s.PrecatenateLength()
}

// TestPrepend tests Slice.Prepend.
func TestPrepend(t *testing.T) {
	s := &slice.Slice[int]{}
	s.Prepend()
}

// TestPrependLength tests Slice.PrependLength.
func TestPrependLength(t *testing.T) {
	s := &slice.Slice[int]{}
	s.PrependLength()
}

// TestPush tests Slice.Push.
func TestPush(t *testing.T) {
	s := &slice.Slice[int]{}
	s.Push()
}

// TestReplace tests Slice.Replace.
func TestReplace(t *testing.T) {
	s := &slice.Slice[int]{}
	s.Replace()
}

// TestReverse tests Slice.Reverse.
func TestReverse(t *testing.T) {
	s := &slice.Slice[int]{}
	s.Reverse()
}

// TestSet tests Slice.Set.
func TestSet(t *testing.T) {
	s := &slice.Slice[int]{}
	s.Set()
}

// TestSlice tests Slice.Slice.
func TestSlice(t *testing.T) {
	s := &slice.Slice[int]{}
	s.Slice()
}

// TestSwap tests Slice.Swap.
func TestSwap(t *testing.T) {
	s := &slice.Slice[int]{}	
	s.Swap()
}

// TestUnshift tests Slice.Unshift.
func TestUnshift(t *testing.T) {
	s := &slice.Slice[int]{}
	s.Unshift()
}
