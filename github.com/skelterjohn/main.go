package main
import (
	"fmt"
	"github.com/skelterjohn/go.matrix"
)
func main() {
	mat := matrix.MakeDenseMatrix([]float64{1, 2, 3, 4, 5, 6}, 2, 3)
	//col,row:=mat.GetSize()
	fmt.Println(mat)
	fmt.Println(matrix.MakeDenseCopy(mat).Transpose())
	fmt.Println(matrix.MakeSparseMatrix(map[int]float64{5: 2.44}, 2, 3))
	fmt.Println("%v %v", mat.Get(1, 1))
}