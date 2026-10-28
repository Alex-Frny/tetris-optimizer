package main

import (
	"fmt"
	"strings"
)

func parseTetrominoes(data string) ([]Tetromino, error) {
	// Accepte les fins de ligne Windows et Unix.
	data = strings.ReplaceAll(data, "\r\n", "\n")
	// Retire uniquement le saut de ligne final, sans masquer les lignes en trop.
	data = strings.TrimSuffix(data, "\n")
	groups := strings.Split(data, "\n\n")
	if len(groups) > 26 {
		return nil, fmt.Errorf("trop de pièces pour les lettres A à Z")
	}

	pieces := make([]Tetromino, 0, len(groups))
	for i, group := range groups {
		blocks, err := validateTetromino(strings.Split(group, "\n"))
		if err != nil {
			return nil, fmt.Errorf("pièce %d : %w", i+1, err)
		}
		pieces = append(pieces, Tetromino{
			Letter: byte('A' + i),
			Blocks: blocks,
		})
	}
	return pieces, nil
}
