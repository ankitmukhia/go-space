package nestedslice
// https://www.boot.dev/lessons/ce8736ca-c203-4cb5-ace0-1c5294f98941
// next quick matrix : https://www.youtube.com/watch?v=JMjbPh1Mjn8&t=40s
// best ::  

func createMatrix(rows, cols int) [][]int {
	matrixResult := [][]int{}

	for i := 0; i < rows; i++ {
		/* end of the loop it will have this 
		* -> [0, 0, 0]
		* -> [0, 1, 2]
		* -> [0, 2, 4]
		*/
		row := []int{}
		for j := 0; j < cols; j++ {
			// calculate i * j
			// the cols value -> put row by row, in each outer loop
			/*
			* 1st(outer) -> [0, 0, 0]
			* 2st(outer) -> [0, 1, 2]
			* 3st(outer) -> [0, 2, 4]
			*/
			row = append(row, i * j)
		}
		matrixResult = append(matrixResult, row)
	}

	return matrixResult
}

