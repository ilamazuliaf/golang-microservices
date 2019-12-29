package utils

import "sort"

func BubbleSort(elements []int) {
	keepRunning := true
	for keepRunning {
		keepRunning = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}
	}
}

// use native go sort function when > 1000
func Sort(els []int) {
	if len(els) < 1000 {
		BubbleSort(els)
		return
	}
	sort.Ints(els)
}
