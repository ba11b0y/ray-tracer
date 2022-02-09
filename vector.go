package main

import (
	"fmt"
	"math"
	"strings"
)

type Vector struct {
	X float64
	Y float64
	Z float64
}

func (v1 *Vector) add(v2 *Vector) {
	v1.X += v2.X
	v1.Y += v2.Y
	v1.Z += v2.Z
}

func (v1 *Vector) mul(v2 Vector) {
	v1.X *= v2.X
	v1.Y *= v2.Y
	v1.Z *= v2.Z
}

func (v1 *Vector) mulConst(c float64) {
	v1.X *= c
	v1.Y *= c
	v1.Y *= c
}

func (v1 Vector) length() float64 {
	return math.Sqrt(math.Pow(v1.X, 2) + math.Pow(v1.Y, 2) + math.Pow(v1.Z, 2))
}

func (v1 Vector) writeColor(f *strings.Builder) {
	var ir, ig, ib int
	ir = int(255.999 * v1.X)
	ig = int(255.999 * v1.Y)
	ib = int(255.999 * v1.Z)

	f.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
}

func Dot(v1 *Vector, v2 *Vector) Vector {
	return Vector{
		X: v1.X * v2.X,
		Y: v1.Y * v2.Y,
		Z: v1.Z * v2.Z,
	}
}

func Cross(v1 *Vector, v2 *Vector) Vector {
	return Vector{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.X,
	}
}
