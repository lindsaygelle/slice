package slice

type slicer []interface{}

func (slicer *slicer) each(fn func(int, interface{})) {
	for i, x := range *slicer {
		fn(i, x)
	}
}

func (slicer *slicer) eachString(fn func(int, string)) {
	slicer.each(func(i int, x interface{}) {
		fn(i, (x.(string)))
	})
}

func (slicer *slicer) push(i interface{}) {
	(*slicer) = (append(*slicer, i))
}

func (slicer *slicer) pushString(s string) {
	slicer.push(s)
}
