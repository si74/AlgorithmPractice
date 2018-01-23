/* matrix multiplier is a package that tests out using separate goroutines to speed up
   matrix calculation. for small N, it is actually slower however ironically */
package matrixmultiplier

import (
	"fmt"
	"sync"
)

// validates that each matrix is proper rectangular matrix
func validate([][]int) [][]int {
	return nil
}

// ResultMatrix is struct with mutex to allow for concurrent calculation
type ResultMatrix struct {
	res [][]int
	sync.Mutex
}

func (r *ResultMatrix) initalize(N int) {
	// initialize empty result array
	res := make([][]int, N)
	for i := 0; i < N; i++ {
		res[i] = make([]int, N)
	}
	r.res = res
}

// MultiplySquare multiplies N-dimensional square matrices
func MultiplySquare(a [][]int, b [][]int, N int) [][]int {
	// initialize empty result array
	res := make([][]int, N)
	for i := 0; i < N; i++ {
		res[i] = make([]int, N)
	}
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
	// initialize empty result array
	result := ResultMatrix{}
	result.initalize(N)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			wg.Add(1)
			go func(i int, j int) {
				defer wg.Done()
				temp := 0
				for k := 0; k < N; k++ {
					fmt.Println(k)
					temp += a[i][k] * b[k][j]
				}
				result.Lock()
				result.res[i][j] = temp
				result.Unlock()
			}(i, j)
		}
	}
	wg.Wait()
	return result.res
}
