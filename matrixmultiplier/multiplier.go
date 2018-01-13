package matrixmultiplier

import "sync"

// validates that each matrix is proper rectangular matrix
func validate([][]int) [][]int {
	return nil
}

// ResultMatrix is struct with mutex to allow for concurrent calculation
type ResultMatrix struct {
	res [][]int
	sync.Mutex
}

// MultiplySquare multiplies N-dimensional square matrices
func MultiplySquare(a [][]int, b [][]int, N int) [][]int {
	var res [][]int
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			res[i][j] = 0
			for k := 0; k < N; k++ {
				res[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return res
}

// MultiplySquareConcurrently multiplies N-dimensional square matrices but leverages waitgroups
func MultiplySquareConcurrently(a [][]int, b [][]int, N int) [][]int {
	var wg sync.WaitGroup
	result := ResultMatrix{}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				temp := 0
				for k := 0; k < N; k++ {
					temp += a[i][k] * b[k][j]
				}
				result.Lock()
				result.res[i][j] = temp
				result.Unlock()
			}()
		}
	}
	wg.Wait()
	return result.res
}
