package basic

type QuickUnionUF struct {
	size    int
	parents []int
	count   int
}

func NewQuickUnionUF(n int) *QuickUnionUF {
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}
	return &QuickUnionUF{n, parents, n}

}

func (uf *QuickUnionUF) validate(n int) bool {
	if n < 0 || n >= uf.size {
		return false
	}
	return true
}

func (uf *QuickUnionUF) Find(p int) int {
	if !uf.validate(p) {
		return -1
	}
	var parent int
	for {
		parent = uf.parents[p]
		if parent == p {
			break
		}
		p = parent
	}
	return parent
}

func (uf *QuickUnionUF) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *QuickUnionUF) Count() int {
	return uf.count
}

func (uf *QuickUnionUF) Union(p, q int) {
	if !uf.validate(p) {
		return
	}
	if !uf.validate(q) {
		return
	}
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	if rootP == rootQ {
		return
	}
	uf.parents[rootQ] = p
}
