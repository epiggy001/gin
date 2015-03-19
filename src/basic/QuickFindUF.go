package basic

type QuickFindUF struct {
	count int
	ids   []int
	size  int
}

func NewQuickFindUF(n int) *QuickFindUF {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = i
	}
	return &QuickFindUF{n, ids, n}
}

func (uf *QuickFindUF) validate(n int) bool {
	if n < 0 || n >= uf.size {
		return false
	}
	return true
}

func (uf *QuickFindUF) Find(p int) int {
	if !uf.validate(p) {
		return -1
	}
	return uf.ids[p]
}

func (uf *QuickFindUF) Connected(p, q int) bool {
	if !uf.validate(p) {
		return false
	}
	if !uf.validate(q) {
		return false
	}
	return uf.ids[p] == uf.ids[q]
}

func (uf *QuickFindUF) Count() int {
	return uf.count
}

func (uf *QuickFindUF) Union(p, q int) {
	if uf.Connected(p, q) {
		return
	}
	ids := uf.ids
	t := ids[q]
	for i := 0; i < uf.size; i++ {
		if ids[i] == t {
			ids[i] = p
		}
	}
	uf.count--
}
