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

// gather gathers nodes with indecies of [l, r), and returns its root node and its depth.
// If the index range is invalid, it returns nil, -2.
// If either one of l==0 or r==s.Len() holds, then depth is 1.
// If both are true, then it's 0.
// If the array is empty, it returns nil, 0.
// If the range is empty i.e. l==r, it returns nil, -1.
// Otherwise, the depth is 2.
// l, r are 0-based index.
// If any splay operation happens after calling this function, the returned node is no longer meaningful.
func (s *Splay) gather(l, r int) (*node, int) {
	if l < 0 || r > s.Len() {
		return nil, -2
	}
	if l == r {
		return nil, -1
	}
	if s.Len() == 0 {
		return nil, 0
	}
	if l == 0 && r == s.Len() {
		return s.root, 0
	} else if l == 0 {
		s.splay(s.getKthNode(r))
		return s.root.l, 1
	} else if r == s.Len() {
		s.splay(s.getKthNode(l - 1))
		return s.root.r, 1
	}
	s.splay(s.getKthNode(r))
	origin := s.root
	s.root = s.root.l
	s.splay(s.getKthNode(l - 1))
	target := s.root.r
	s.root = origin
	return target, 2
}
