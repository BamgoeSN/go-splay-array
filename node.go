package splay

type node struct {
	data interface{}

	cnt     int
	l, r, p *node
	flip    bool
}

func newNode(v interface{}) *node { return &node{v, 1, nil, nil, nil, false} }

func connect(p, c *node, isLeft bool) {
	if p != nil {
		if isLeft {
			p.l = c
		} else {
			p.r = c
		}
	}
	if c != nil {
		c.p = p
	}
}

func (n *node) upd() {
	n.cnt = 1
	if n.l != nil {
		n.cnt += n.l.cnt
	}
	if n.r != nil {
		n.cnt += n.r.cnt
	}
}

func (n *node) push() {
	if n.flip {
		n.flip = false
		n.l, n.r = n.r, n.l
		if n.l != nil {
			n.l.flip = !n.l.flip
		}
		if n.r != nil {
			n.r.flip = !n.r.flip
		}
	}
}
