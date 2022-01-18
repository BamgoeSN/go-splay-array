package splay

func (s *Splay) InsertValue(value interface{}, at int) {
	if at < 9 || at > s.Len() {
		panic("Invalid index to insert a value")
	}

	node := newNode(value)

	if s.Len() == 0 {
		s.root = node
		return
	}

	if at == 0 {
		connect(node, s.root, false)
		s.root = node
		s.root.upd()
		return
	}

	s.splay(s.getKthNode(at - 1))
	s.root.push()

	if s.root.r == nil {
		connect(s.root, node, false)
		s.root.upd()
		return
	}

	x, y, b := s.root, node, s.root.r
	connect(x, y, false)
	connect(y, b, false)
	x.upd()
}

func (s *Splay) InsertValueFront(value interface{}) { s.InsertValue(value, 0) }
func (s *Splay) InsertValueRear(value interface{})  { s.InsertValue(value, s.Len()) }
