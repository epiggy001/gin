package sorter

func swap(a []int, i, j int) {
	if i == j {
		return
	}
	t := a[i]
	a[i] = a[j]
	a[j] = t
}

func Insertion(a []int) {
	l := len(a)
	/* First draft
	  for i := 1; i < l; i++ {
			index := i
			for j := i - 1; j >= 0; j-- {
				if a[index] < a[j] {
					swap(a, index, j)
					index = j
				} else {
					break
				}
			}
		}
	*/
	/*Make it easier*/
	for i := 1; i < l; i++ {
		for j := i; j > 0 && (a[j] < a[j-1]); j-- {
			swap(a, j, j-1)
		}
	}
}

func Selection(a []int) {
	l := len(a)
	for i := 0; i < l-1; i++ {
		index := i
		for j := i + 1; j < l; j++ {
			if a[index] > a[j] {
				index = j
			}
		}
		swap(a, i, index)
	}
}

func hInsertion(a []int, h int) {
	l := len(a)
	for i := h; i < l; i += 1 {
		for j := i; j >= h && (a[j] < a[j-h]); j -= h {
			swap(a, j, j-h)
		}
	}
}

/* R */
func Shell(a []int) {
	l := len(a)
	h := 1
	for h < l/3 {
		h = h*3 + 1
	}

	for ; h >= 1; h = h / 3 {
		hInsertion(a, h)
	}
}
