package splay

type Splay struct {
	root *node
}

func EmptySplay() *Splay {
	return &Splay{nil}
}

// Len returns the number of elements in the splay array.
func (s *Splay) Len() int {
	if s.root == nil {
		return 0
	}
	return s.root.cnt
}

// CopyToSlice returns a slice converted from the splay array. The slice returned is not synced with the splay array;
// i.e. even if the splay array is updated later on, the returned slice remains the same.
func (s *Splay) CopyToSlice() []interface{} {
	if s.root == nil {
		return nil
	}
	arr := make([]interface{}, 0, s.root.cnt)
	s.toSliceHelper(&arr, s.root)
	return arr
}
