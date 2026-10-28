package main

import "fmt"

func validateTetromino(lines []string) ([4]Point, error) {
	var blocks [4]Point
	if len(lines) != 4 {
		return blocks, fmt.Errorf("il faut quatre lignes")
	}

	count := 0
	minX, minY := 4, 4
	for y, line := range lines {
		if len(line) != 4 {
			return blocks, fmt.Errorf("il faut quatre caractères par ligne")
		}
		for x := 0; x < 4; x++ {
			switch line[x] {
			case '.':
			case '#':
				if count == 4 {
					return blocks, fmt.Errorf("trop de cases occupées")
				}
				blocks[count] = Point{X: x, Y: y}
				count++
				if x < minX {
					minX = x
				}
				if y < minY {
					minY = y
				}
			default:
				return blocks, fmt.Errorf("caractère interdit")
			}
		}
	}
	if count != 4 {
		return blocks, fmt.Errorf("il faut exactement quatre cases occupées")
	}

	// Parcourt les cases accessibles par un côté depuis la première case.
	visited := [4]bool{true}
	queue := []int{0}
	for head := 0; head < len(queue); head++ {
		current := blocks[queue[head]]
		for i, block := range blocks {
			dx, dy := current.X-block.X, current.Y-block.Y
			if !visited[i] && (dx*dx+dy*dy == 1) {
				visited[i] = true
				queue = append(queue, i)
			}
		}
	}
	if len(queue) != 4 {
		return blocks, fmt.Errorf("les cases doivent être reliées par leurs côtés")
	}

	// Ramène la pièce en haut à gauche pour faciliter son futur placement.
	for i := range blocks {
		blocks[i].X -= minX
		blocks[i].Y -= minY
	}
	return blocks, nil
}
