package main

import (
	"fmt"
	"math"
)

type lingkaran struct {
	jari float64
}

type bentuk interface {
	Area() float64
}

func (L lingkaran) Area() float64 {
	return math.Pi * (L.jari * L.jari)
}

func getBentuk(b bentuk) float64 {
	return b.Area()
}

func main() {
	circle := lingkaran{jari: 4.5}
	fmt.Printf("Luas Lingkaran adalah : %v", getBentuk(circle))
}
