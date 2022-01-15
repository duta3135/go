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
	var duta Vector
	duta.v = 5
	duta.w = 1
	duta.x = 4
	duta.y = 5
	duta.z = 4
	var jespan Vector
	jespan.v = 1
	jespan.w = 2
	jespan.x = 5
	jespan.y = 2
	jespan.z = 4
	dotProduct := Vector{duta.v, duta.w, duta.x, duta.y, duta.z}.Dot(Vector{jespan.v, jespan.w, jespan.x, jespan.y, jespan.z})
	amplitude1 := Vector{duta.v, duta.w, duta.x, duta.y, duta.z}.Amplitude()
	amplitude2 := Vector{jespan.v, jespan.w, jespan.x, jespan.y, jespan.z}.Amplitude()
	var cosineSim float64 = dotProduct / (amplitude1 * amplitude2)
	fmt.Printf("kesamaan cosinus antara jespan dan duta: %f", cosineSim)
}
