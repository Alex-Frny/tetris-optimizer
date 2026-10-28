package main

import (
	"strings"
	"testing"
)

func TestParseTetrominoes(t *testing.T) {
	square := "....\n....\n..##\n..##"
	for _, ending := range []string{"", "\n"} {
		for _, newline := range []string{"\n", "\r\n"} {
			input := strings.ReplaceAll(square+"\n\n"+square+ending, "\n", newline)
			pieces, err := parseTetrominoes(input)
			if err != nil {
				t.Fatal(err)
			}
			want := [4]Point{{0, 0}, {1, 0}, {0, 1}, {1, 1}}
			if len(pieces) != 2 || pieces[0].Letter != 'A' || pieces[1].Letter != 'B' || pieces[0].Blocks != want {
				t.Fatalf("résultat inattendu : %+v", pieces)
			}
		}
	}
}

func TestRejectInvalidInput(t *testing.T) {
	square := "##..\n##..\n....\n...."
	cases := map[string]string{
		"empty":               "",
		"missing row":         "####\n....\n....",
		"long row":            "####.\n....\n....\n....",
		"invalid character":   "##x.\n##..\n....\n....",
		"too few blocks":      "###.\n....\n....\n....",
		"too many blocks":     "####\n#...\n....\n....",
		"disconnected":        "##..\n....\n..##\n....",
		"diagonal":            "#...\n.#..\n..#.\n...#",
		"extra separator":     square + "\n\n\n" + square,
		"missing separator":   square + "\n" + square,
		"trailing blank line": square + "\n\n",
		"leading blank line":  "\n" + square,
		"too many pieces":     strings.Repeat(square+"\n\n", 26) + square,
	}
	for name, input := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := parseTetrominoes(input); err == nil {
				t.Fatal("le fichier aurait dû être rejeté")
			}
		})
	}
}

func TestAllConnectedShapes(t *testing.T) {
	// Les 19 orientations possibles, sans rotation effectuée par le parser.
	shapes := []string{
		"####", "#/#/#/#", "##/##",
		"###/.#", ".#/###", "#./##/#", ".#/##/.#",
		".##/##", "#./##/.#", "##/.##", ".#/##/#",
		"#/#/##", "###/#", "##/.#/.#", "..#/###",
		".#/.#/##", "#/###", "##/#/#", "###/..#",
	}
	for _, shape := range shapes {
		rows := strings.Split(shape, "/")
		for i := range rows {
			rows[i] += strings.Repeat(".", 4-len(rows[i]))
		}
		for len(rows) < 4 {
			rows = append(rows, "....")
		}
		if _, err := parseTetrominoes(strings.Join(rows, "\n")); err != nil {
			t.Errorf("forme %q rejetée : %v", shape, err)
		}
	}
}
