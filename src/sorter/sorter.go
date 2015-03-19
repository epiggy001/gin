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

var counter int = 0

func mergeSort(a []int, i, j int) {
	counter++
	if i >= j {
		return
	}

	m := (j + i) / 2
	mergeSort(a, i, m)
	mergeSort(a, m+1, j)
	b := make([]int, j-i+1)
	index1 := i
	index2 := m + 1
	/* first draft
	n := 0
	for {
		if a[index1] < a[index2] {
			b[n] = a[index1]
			index1++
		} else if a[index1] >= a[index2] {
			b[n] = a[index2]
			index2++
		}

		n++
		if (index1 > m) || (index2 > j) {
			break
		}
	}

	for ; index1 <= m; index1++ {
		b[n] = a[index1]
		n++
	}

	for ; index2 <= j; index2++ {
		b[n] = a[index2]
		n++
	}*/

	for n := i; n <= j; n++ {
		if index1 > m {
			b[n-i] = a[index2]
			index2++
		} else if index2 > j {
			b[n-i] = a[index1]
			index1++
		} else if a[index1] < a[index2] {
			b[n-i] = a[index1]
			index1++
		} else {
			b[n-i] = a[index2]
			index2++
		}
	}

	for n := 0; n < j-i+1; n++ {
		a[i+n] = b[n]
	}
}

func Merge(a []int) {
	l := len(a)
	mergeSort(a, 0, l-1)
}
