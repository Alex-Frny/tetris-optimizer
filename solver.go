package main

// solve attend des pièces validées et normalisées par parseTetrominoes.
func solve(pieces []Tetromino) [][]byte {
	size := 1
	for size*size < len(pieces)*4 {
		size++
	}

	for {
		board := makeBoard(size)
		if placePieces(board, pieces, 0) {
			return board
		}
		size++
	}
}

func makeBoard(size int) [][]byte {
	board := make([][]byte, size)
	for y := range board {
		board[y] = make([]byte, size)
		for x := range board[y] {
			board[y][x] = '.'
		}
	}
	return board
}

func placePieces(board [][]byte, pieces []Tetromino, index int) bool {
	if index == len(pieces) {
		return true
	}

	piece := pieces[index]
	for y := range board {
		for x := range board[y] {
			if !canPlace(board, piece, x, y) {
				continue
			}
			writePiece(board, piece, x, y, piece.Letter)
			if placePieces(board, pieces, index+1) {
				return true
			}
			// Cette position bloque la suite : on efface et on réessaie ailleurs.
			writePiece(board, piece, x, y, '.')
		}
	}
	return false
}

func canPlace(board [][]byte, piece Tetromino, x, y int) bool {
	for _, block := range piece.Blocks {
		px, py := x+block.X, y+block.Y
		if py >= len(board) || px >= len(board[py]) || board[py][px] != '.' {
			return false
		}
	}
	return true
}

func writePiece(board [][]byte, piece Tetromino, x, y int, letter byte) {
	for _, block := range piece.Blocks {
		board[y+block.Y][x+block.X] = letter
	}
}
