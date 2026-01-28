package main

import (
	"fmt"
	"math"
)

// https://pkg.go.dev/math

func main() {
	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1, 2))
	fmt.Println(math.Abs(-1))
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Sqrt(4))
	fmt.Println(math.Ceil(1.1))
	fmt.Println(math.Floor(1.9))
	fmt.Println(math.Round(1.4))
	fmt.Println(math.RoundToEven(1.5))
	fmt.Println(math.Trunc(1.9))
	fmt.Println(math.Mod(1, 2))
	fmt.Println(math.Log(1))
	fmt.Println(math.Log10(1))
	fmt.Println(math.Log2(1))
	fmt.Println(math.Exp(1))
	fmt.Println(math.Hypot(1, 2))
	fmt.Println(math.Sin(1))
	fmt.Println(math.Cos(1))
	fmt.Println(math.Tan(1))
	fmt.Println(math.Asin(1))
	fmt.Println(math.Acos(1))
	fmt.Println(math.Atan(1))
	fmt.Println(math.Atan2(1, 2))
	fmt.Println(math.Pi)
	fmt.Println(math.E)
	fmt.Println(math.MaxInt)
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MaxUint8)
	fmt.Println(math.MaxUint16)
	fmt.Println(math.MaxUint32)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.SmallestNonzeroFloat32)
	fmt.Println(math.SmallestNonzeroFloat64)
	fmt.Println(math.NaN())
	fmt.Println(math.IsNaN(1))
	fmt.Println(math.IsInf(1, 1))
}
