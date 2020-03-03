package main

import (
	"github.com/emirpasic/gods/maps/treemap"
)

func segment(x int, arr []int) int {
	m := treemap.NewWithIntComparator()

	for i := 0; i < x; i++ {
		v, found := m.Get(arr[i])
		if found {
			m.Put(arr[i], int(v)+1)
		} else {
			m.Put(arr[i], 1)
		}
	}

	maxOfMin, _ := m.Min()

	for i := x; i <= len(arr)-x; i++ {
		f := i + x
		b := i - x

		v, found := m.Get(arr[b])
		if found {
			m.Remove(arr[b])
			if v > 1 {
				m.Put(arr[b], v-1)
			}
		}

		v, found = m.Get(arr[f])
		if found {
			m.Put(arr[f], v+1)
		} else {
			m.Put(arr[f], 1)
		}

		minKey, _ := m.Min()
		if minKey > maxOfMin {
			maxOfMin = minKey
		}
	}

	return maxOfMin
}

func main() {
	arr := []int{1, 2, 3, 1, 2}
	ret := segment(2, arr)
}
