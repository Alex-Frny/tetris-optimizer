package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	cases := []struct {
		name  string
		input string
		size  int
	}{
		{"square", "##..\n##..\n....\n....", 2},
		{"vertical bar", "#...\n#...\n#...\n#...", 4},
		{"sample", "#...\n#...\n#...\n#...\n\n....\n....\n..##\n..##", 4},
		{"two squares cannot fit in three", "##..\n##..\n....\n....\n\n##..\n##..\n....\n....", 4},
		{"backtracking", ".##.\n##..\n....\n....\n\n##..\n.##.\n....\n....", 4},
		{"assignment sample", "...#\n...#\n...#\n...#\n\n....\n....\n....\n####\n\n.###\n...#\n....\n....\n\n....\n..##\n.##.\n....\n\n....\n.##.\n.##.\n....\n\n....\n....\n##..\n.##.\n\n##..\n.#..\n.#..\n....\n\n....\n###.\n.#..\n....", 6},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			pieces, err := parseTetrominoes(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			board := solve(pieces)
			if len(board) != tc.size {
				t.Fatalf("taille %d, attendu %d", len(board), tc.size)
			}
			for _, row := range board {
				if len(row) != tc.size {
					t.Fatal("grille non carrée")
				}
				for _, cell := range row {
					if cell != '.' && (cell < 'A' || int(cell-'A') >= len(pieces)) {
						t.Fatalf("case inconnue : %q", cell)
					}
				}
			}
			for _, piece := range pieces {
				var cells []Point
				minX, minY := tc.size, tc.size
				for y, row := range board {
					for x, cell := range row {
						if cell == piece.Letter {
							cells = append(cells, Point{x, y})
							if x < minX {
								minX = x
							}
							if y < minY {
								minY = y
							}
						}
					}
				}
				if len(cells) != 4 {
					t.Fatalf("pièce %c : %d cases", piece.Letter, len(cells))
				}
				for i, cell := range cells {
					if (Point{cell.X - minX, cell.Y - minY}) != piece.Blocks[i] {
						t.Fatalf("pièce %c déformée", piece.Letter)
					}
				}
			}
		})
	}
}

func TestFailedSearchRestoresBoard(t *testing.T) {
	pieces, err := parseTetrominoes("##..\n##..\n....\n....\n\n##..\n##..\n....\n....")
	if err != nil {
		t.Fatal(err)
	}
	board := makeBoard(3)
	if placePieces(board, pieces, 0) {
		t.Fatal("deux carrés ne peuvent pas tenir dans une grille 3x3")
	}
	for _, row := range board {
		if string(row) != strings.Repeat(".", 3) {
			t.Fatal("le retour arrière a laissé des cases occupées")
		}
	}
}
