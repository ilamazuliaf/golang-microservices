package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubleSortWorstCase(t *testing.T) {
	// inisialisasi
	els := []int{9,8,7,6,5}
	//eksekusi
	Sort(els)
	// validasi
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

func TestBubleSortBestCase(t *testing.T) {
	// inisialisasi
	els := []int{5,6,7,8,9}
	//eksekusi
	Sort(els)
	// validasi
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

func getElement(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n-1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return  result
}

func TestGetElements(t *testing.T) {
	els := getElement(5)

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 4, els[0])
	assert.EqualValues(t, 3, els[1])
	assert.EqualValues(t, 2, els[2])
	assert.EqualValues(t, 1, els[3])
	assert.EqualValues(t, 0, els[4])
}

func BenchmarkBubleSort10(b *testing.B) {
	els := getElement(10)
	for i :=0; i < b.N; i++ {
		Sort(els)
	}
}

// Native go sort function
func BenchmarkSort10(b *testing.B) {
	els := getElement(10)
	for i :=0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkBubleSort1000(b *testing.B) {
	els := getElement(1000)
	for i :=0; i < b.N; i++ {
		Sort(els)
	}
}

// Native go sort function
func BenchmarkSort1000(b *testing.B) {
	els := getElement(1000)
	for i :=0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkBubleSort50000(b *testing.B) {
	els := getElement(50000)
	for i :=0; i < b.N; i++ {
		Sort(els)
	}
}

// Native go sort function
func BenchmarkSort50000(b *testing.B) {
	els := getElement(50000)
	for i :=0; i < b.N; i++ {
		Sort(els)
	}
}