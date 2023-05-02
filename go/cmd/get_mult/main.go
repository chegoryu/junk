package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d%d", &a, &b)
	fmt.Println("Mult:", GetMult(a, b))
}
