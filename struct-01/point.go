package main

import "image/color"

type Point struct {
	X, Y int
}

type ColoredPoint struct {
	*Point
	Color color.RGBA
}