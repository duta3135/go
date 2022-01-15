package main

import (
	"fmt"
	"math"
)

type Vector struct {
	v, w, x, y, z float64
}

func (vectorA Vector) Dot(vectorB Vector) float64 {
	var dotProd float64 = vectorA.v*vectorB.v + vectorA.w*vectorB.w + vectorA.x*vectorB.x + vectorA.y*vectorB.y + vectorA.z*vectorB.z
	return dotProd
}

func (a Vector) Amplitude() float64 {
	return math.Sqrt(a.Dot(a))
}
func main() {
	var lisia Vector
	lisia.v = 1
	lisia.w = 1
	lisia.x = 1
	lisia.y = 1
	lisia.z = 1
	var leri Vector
	leri.v = 5
	leri.w = 5
	leri.x = 5
	leri.y = 5
	leri.z = 5
	cosineSim := (Vector{lisia.v, lisia.w, lisia.x, lisia.y, lisia.z}.Dot(Vector{leri.v, leri.w, leri.x, leri.y, leri.z})) / (Vector{lisia.v, lisia.w, lisia.x, lisia.y, lisia.z}.Amplitude() * Vector{leri.v, leri.w, leri.x, leri.y, leri.z}.Amplitude())
	fmt.Printf("lisia and leri's cosine similarity: %f", cosineSim)
}
