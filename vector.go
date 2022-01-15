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
	var ine Vector
	ine.v = 1
	ine.w = 1
	ine.x = 2
	ine.y = 4
	ine.z = 1
	var petrik Vector
	petrik.v = 5
	petrik.w = 1
	petrik.x = 1
	petrik.y = 1
	petrik.z = 1
	dotProduct := Vector{ine.v, ine.w, ine.x, ine.y, ine.z}.Dot(Vector{petrik.v, petrik.w, petrik.x, petrik.y, petrik.z})
	amplitude1 := Vector{ine.v, ine.w, ine.x, ine.y, ine.z}.Amplitude()
	amplitude2 := Vector{petrik.v, petrik.w, petrik.x, petrik.y, petrik.z}.Amplitude()
	var cosineSim float64 = dotProduct / (amplitude1 * amplitude2)
	fmt.Printf("kesamaan cosinus antara petrik dan ine: %f", cosineSim)
}
