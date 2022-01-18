package splay

func (s *Splay) Flip(l, r int) {
	if l > r || l < 0 || r > s.Len() {
		panic("Invalid range")
	}
	if l == r {
		return
	}
	ptr, _ := s.gather(l, r)
	ptr.flip = !ptr.flip
}

func (s *Splay) CopyRangeToSlice(l, r int) []interface{} {
	if l > r || l < 0 || r > s.Len() {
		panic("Invalid range")
	}
	if l == r {
		return nil
	}
	ptr, _ := s.gather(l, r)
	arr := make([]interface{}, 0, ptr.cnt)
	s.toSliceHelper(&arr, ptr)
	return arr
}

func (s *Splay) TakeOutRange(l, r int) *Splay {
	if l > r || l < 0 || r > s.Len() {
		panic("Invalid range")
	}
	if l == r {
		return EmptySplay()
	}

	x, d := s.gather(l, r)
	x.push()
	if d == 0 {
		ans := &Splay{s.root}
		s.root = nil
		return ans
	}

	p := x.p
	if p.l == x {
		p.l = nil
	} else {
		p.r = nil
	}
	x.p = nil

	for p != nil {
		p.upd()
		p = p.p
	}

	return &Splay{x}
}

func (s *Splay) InsertSplay(pattern *Splay, at int) {
	defer func() { pattern.root = nil }()

	if at < 0 || at > s.Len() {
		panic("Invalid index to insert a splay")
	}

	if s.Len() == 0 {
		s.root = pattern.root
		return
	}

	if pattern.Len() == 0 {
		return
	}
	pattern.root.push()

	if at == 0 {
		pattern.splay(pattern.getRightMost())
		connect(pattern.root, s.root, false)
		s.root = pattern.root
		s.root.upd()
		return
	}

	s.splay(s.getKthNode(at - 1))
	s.root.push()

	if s.root.r == nil {
		connect(s.root, pattern.root, false)
		s.root.upd()
		return
	}

	pattern.splay(pattern.getRightMost())

	x, y, b := s.root, pattern.root, s.root.r
	connect(x, y, false)
	connect(y, b, false)
	x.upd()
}
