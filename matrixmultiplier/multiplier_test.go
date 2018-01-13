package matrixmultiplier

import (
	"reflect"
	"testing"
	"time"
)

func Test_MatrixSquare(t *testing.T) {

	a := [][]int{{1, 1}, {2, 3}}
	b := [][]int{{1, 1}, {2, 4}}
	expectedRes := [][]int{{3, 5}, {8, 14}}

	start := time.Now()
	res := MultiplySquare(a, b, 2)
	end := time.Now()

	t.Logf("time elapsed: %v", end.Sub(start))
	if !reflect.DeepEqual(res, expectedRes) {
		t.Logf("want: %v", a)
		t.Logf("got: %v", b)
		t.Fatal("unexpected result matrix")
	}
}

func Test_MatrixSquareConcurrently(t *testing.T) {}
