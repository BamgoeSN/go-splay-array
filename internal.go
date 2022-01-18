package splay

// rotate panics if x == s.root
func (s *Splay) rotate(x *node) {
	x.p.push()
	x.push()

	if s.root == x.p {
		s.root = x
	}

	p := x.p
	pp := x.p.p
	if pp != nil {
		connect(pp, x, pp.l == p)
	} else {
		x.p = nil
	}

	if p.l == x {
		b := x.r
		connect(p, b, true)
		connect(x, p, false)
	} else {
		b := x.l
		connect(p, b, false)
		connect(x, p, true)
	}

	p.upd()
	x.upd()
}

func (s *Splay) splay(x *node) {
	for {
		if x == s.root {
			return
		} else if x.p == s.root {
			s.rotate(x)
			return
		}

		p, g := x.p, x.p.p
		if (p.l == x) == (g.l == p) {
			s.rotate(p)
			s.rotate(x)
		} else {
			s.rotate(x)
			s.rotate(x)
		}
	}
}

func (s *Splay) getKthNode(k int) *node {
	if k >= s.Len() || k < 0 || s.Len() == 0 {
		return nil
	}
	ptr := s.root
	for {
		ptr.push()
		left := 0
		if ptr.l != nil {
			left = ptr.l.cnt
		}
		if left == k {
			break
		} else if left > k {
			ptr = ptr.l
		} else {
			k -= left + 1
			ptr = ptr.r
		}
	}
	ptr.push()
	return ptr
}
