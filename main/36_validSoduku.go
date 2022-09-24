package main

// func main() {
// 	b1 := [][]byte{
// 		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
// 		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
// 		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
// 		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
// 		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
// 		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
// 		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
// 		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
// 		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
// 	}
// 	fmt.Println(isValidSudoku(b1))
// }

func isValidSudoku(board [][]byte) bool {
	for i, row := range board {
		if !isValidRow(row) {
			return false
		}

		if !isValidColumn(board, i) {
			return false
		}
	}

	squares := [][]int{
		{0, 2},
		{3, 5},
		{6, 8},
	}

	for _, squareX := range squares {
		for _, squareY := range squares {
			if !isValidSquare(board, squareX, squareY) {
				return false
			}
		}
	}

	return true
}

func isValidRow(row []byte) bool {
	numMap := make(map[byte]bool)

	for _, item := range row {
		if item == '.' {
			continue
		}
		_, ok := numMap[item]
		if ok {
			return false
		} else {
			numMap[item] = true
		}
	}
	return true
}

func isValidColumn(board [][]byte, column int) bool {
	numMap := make(map[byte]bool)

	for _, row := range board {
		if row[column] == '.' {
			continue
		}
		_, ok := numMap[row[column]]
		if ok {
			return false
		} else {
			numMap[row[column]] = true
		}
	}
	return true
}

func isValidSquare(board [][]byte, x []int, y []int) bool {
	numMap := make(map[byte]bool)

	for xIndex := x[0]; xIndex <= x[1]; xIndex++ {
		for yIndex := y[0]; yIndex <= y[1]; yIndex++ {
			val := board[yIndex][xIndex]
			if val == '.' {
				continue
			}
			_, ok := numMap[val]
			if ok {
				return false
			} else {
				numMap[val] = true
			}
		}
	}
	return true
}
