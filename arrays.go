package arrays

func IsEqual(a, b []int) bool {
	size := len(a)
	if size != len(b) {
		return false
	}
	for i := 0; i < size; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func IsEqual16(a, b []uint16) bool {
	size := len(a)
	if size != len(b) {
		return false
	}
	for i := 0; i < size; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func IntersectEx(r, a, b []int) []int {
	asize := len(a)
	bsize := len(b)
	i := 0
	j := 0
	r = r[:0]
	for i < asize && j < bsize {
		if a[i] < b[j] {
			i++
			continue
		}
		if a[i] > b[j] {
			j++
			continue
		}
		r = append(r, a[i])
		i++
		j++
	}
	return r
}

func Intersect(a, b []int) []int {
	size := len(a)
	if size > len(b) {
		size = len(b)
	}
	r := make([]int, 0, size)
	return IntersectEx(r, a, b)
}

func UnionAllEx(r, a, b []int) []int {
	asize := len(a)
	bsize := len(b)
	i := 0
	j := 0
	r = r[:0]
	for i < asize && j < bsize {
		if a[i] < b[j] {
			r = append(r, a[i])
			i++
			continue
		}
		if a[i] > b[j] {
			r = append(r, b[j])
			j++
			continue
		}
		r = append(r, a[i], b[j])
		i++
		j++
	}
	for i < asize {
		r = append(r, a[i])
		i++
	}
	for j < bsize {
		r = append(r, b[j])
		j++
	}
	return r
}

func UnionAll(a, b []int) []int {
	size := len(a) + len(b)
	r := make([]int, 0, size)
	return UnionAllEx(r, a, b)
}

func UnionEx(r, a, b []int) []int {
	asize := len(a)
	bsize := len(b)
	i := 0
	j := 0
	r = r[:0]
	for i < asize && j < bsize {
		v := a[i]
		if v > b[j] {
			v = b[j]
		}
		for i < asize {
			if a[i] > v {
				break
			}
			i++
		}
		for j < bsize {
			if b[j] > v {
				break
			}
			j++
		}
		r = append(r, v)
	}
	for i < asize {
		v := a[i]
		for i < asize {
			if a[i] > v {
				break
			}
			i++
		}
		r = append(r, v)
	}
	for j < bsize {
		v := b[j]
		for j < bsize {
			if b[j] > v {
				break
			}
			j++
		}
		r = append(r, v)
	}
	return r

}

func Union(a, b []int) []int {
	size := len(a) + len(b)
	r := make([]int, 0, size)
	return UnionEx(r, a, b)
}

func Distinct(a []int) []int {
	i := 0
	j := 0
	size := len(a)
	for i < size {
		n := i + 1
		if n < size {
			if a[i] == a[n] {
				i++
				continue
			}
		}
		a[j] = a[i]
		i++
		j++
	}
	return a[:j]
}

func ExceptEx(r, a, b []int) []int {
	i := 0
	j := 0
	asize := len(a)
	bsize := len(b)
	r = r[:0]
	for i < asize {
		uniq := true
		for j < bsize {
			if b[j] < a[i] {
				j++
				continue
			}
			if b[j] == a[i] {
				uniq = false
			}
			break
		}
		if uniq {
			r = append(r, a[i])
		}
		i++
	}
	return r
}

func Except(a, b []int) []int {
	size := len(a)
	r := make([]int, 0, size)
	return ExceptEx(r, a, b)
}

func ExceptInpl(a, b []int) []int {
	i := 0
	j := 0
	k := 0
	asize := len(a)
	bsize := len(b)
	for i < asize {
		uniq := true
		for j < bsize {
			if b[j] < a[i] {
				j++
				continue
			}
			if b[j] == a[i] {
				uniq = false
			}
			break
		}
		a[k] = a[i]
		if uniq {
			k++
		}
		i++
	}
	return a[:k]
}

func Compact32Ex(r []uint16, a []uint32) []uint16 {
	i := 0
	j := uint16(0)
	asize := len(a)
	r = append(r[:0], 0, 0)
	for i < asize {
		l := uint16(a[i])
		h := uint16(a[i] >> 16)
		if r[j] != h {
			j += r[j+1] + 2
			r = append(r, h, 0)
		}
		r[j+1]++
		r = append(r, l)
		i++
	}
	return r
}

func Compact32(a []uint32) []uint16 {
	size := len(a)
	r := make([]uint16, 0, (size >> 2))
	return Compact32Ex(r, a)
}
