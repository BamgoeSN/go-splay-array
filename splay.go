package splay

type Splay struct {
	root *node
}

func (s *Splay) Len() int {
	if s.root == nil {
		return 0
	}
	return s.root.cnt
}
