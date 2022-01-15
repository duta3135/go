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
	return math.Sqrt(a.v*a.v + a.w*a.w + a.x*a.x + a.y*a.y + a.z*a.z)
}
func main() {
	var lisia Vector
	lisia.v = 2
	lisia.w = 4
	lisia.x = 4
	lisia.y = 5
	lisia.z = 5
	var leri Vector
	leri.v = 4
	leri.w = 5
	leri.x = 3
	leri.y = 5
	leri.z = 4
	dotProduct := Vector{lisia.v, lisia.w, lisia.x, lisia.y, lisia.z}.Dot(Vector{leri.v, leri.w, leri.x, leri.y, leri.z})
	amplitude1 := Vector{lisia.v, lisia.w, lisia.x, lisia.y, lisia.z}.Amplitude()
	amplitude2 := Vector{leri.v, leri.w, leri.x, leri.y, leri.z}.Amplitude()
	var cosineSim float64 = dotProduct / (amplitude1 * amplitude2)
	fmt.Printf("kesamaan cosinus antara leri dan lisia: %f", cosineSim)
}
