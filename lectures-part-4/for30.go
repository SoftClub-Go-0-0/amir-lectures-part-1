package main

import "fmt"
import "math"

func main() {

	var n float64
	var a float64
	var b float64

	fmt.Scan(&n,&a,&b)
	h:=(b-a)/n
	for i:=a;i<=b;i+=h{
		fmt.Println(1-math.Sin(i))
	}

}