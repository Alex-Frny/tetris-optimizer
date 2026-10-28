package main

type Point struct {
	X int
	Y int
}

type Tetromino struct {
	Letter byte
	Blocks [4]Point
}
