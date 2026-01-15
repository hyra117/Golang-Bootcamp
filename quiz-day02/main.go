package main

import "fmt"

const (
	size5 = 5
	size8 = 8
)

func main() {
	matrixDay2()
}

func matrixDay2() {

	fmt.Println()
	displayMatrix(matrix6())

	fmt.Println()
	displayMatrix(matrix7())

	fmt.Println()
	n := 7
	displayMatrixNo8(matrix8(n))

	fmt.Println()
	displayMatrix(matrix9())

	fmt.Println()
	matrix10()
}

func newMatrix(r, c int) [][]int {
	m := make([][]int, r)
	for i := range m {
		m[i] = make([]int, c)
	}
	return m
}

// No. 6
func matrix6() [][]int {
	matrix := newMatrix(size5, size5)
	for i := 0; i < size5; i++ {
		for j := 0; j < size5; j++ {
			if i == j {
				matrix[i][j] = 1
			} else if j > i {
				matrix[i][j] = 10
			} else {
				matrix[i][j] = 20
			}
		}
	}
	return matrix
}

// No. 7
func matrix7() [][]int {
	matrix := newMatrix(size5, size5)
	for i := 0; i < size5; i++ {
		for j := 0; j < size5; j++ {
			if i == j {
				matrix[i][j] = size5 - i
			} else if j > i {
				matrix[i][j] = 20
			} else {
				matrix[i][j] = 10
			}
		}
	}
	return matrix
}

// No. 8
func matrix8(n int) [][]int {
	matrix := newMatrix(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == n-1 || j == n-1 {
				matrix[i][j] = i + j
			}
		}
	}
	return matrix
}

func displayMatrixNo8(matrix [][]int) {
	n := len(matrix)
	for i := range matrix {
		for j := range matrix[i] {
			if i > 0 && i < n-1 && j > 0 && j < n-1 {
				fmt.Printf("    ")
			} else {
				fmt.Printf("%4d", matrix[i][j])
			}
		}
		fmt.Println()
	}
}

// No. 9
func matrix9() [][]int {
	matrix := newMatrix(size8, size8)

	for i := 0; i < size8-1; i++ {
		for j := 0; j < size8-1; j++ {
			val := i + j
			matrix[i][j] = val
			matrix[i][size8-1] += val
			matrix[size8-1][j] += val
			matrix[size8-1][size8-1] += val
		}
	}
	return matrix
}

func displayMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, v := range row {
			fmt.Printf("%4d", v)
		}
		fmt.Println()
	}
}

// No. 10
func matrix10() {
	jawaban := [][]string{
		{"A", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"D", "B", "A", "B", "C", "A", "E", "E", "A", "D"},
		{"E", "D", "D", "A", "C", "B", "E", "E", "A", "D"},
		{"C", "B", "A", "E", "D", "C", "E", "E", "A", "D"},
		{"A", "B", "D", "C", "C", "D", "E", "E", "A", "D"},
		{"B", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
		{"B", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"E", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
	}

	kunci := []string{"D", "B", "D", "C", "C", "D", "A", "E", "A", "D"}

	for i := range jawaban {
		benar := 0
		for j := range kunci {
			if jawaban[i][j] == kunci[j] {
				benar++
			}
		}
		fmt.Printf("Jawaban Siswa %d yang benar : %d\n", i, benar)
	}
}
